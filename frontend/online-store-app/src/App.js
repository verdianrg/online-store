import React, { useState } from 'react';
import { BrowserRouter as Router, Route, Routes, Link, Navigate } from 'react-router-dom';
import ProductList from './components/ProductList';
import Cart from './components/Cart';
import Summary from './components/Summary';
import History from './components/History';
import Login from './components/Login'; // Import the Login component here
import './App.css';

const App = () => {
  const [cart, setCart] = useState([]);
  const [purchases, setPurchases] = useState([]);
  const [authenticated, setAuthenticated] = useState(false);

  const token = localStorage.getItem('token')

  const handleLogin = () => {
    // Handle authentication logic here (e.g., API key and JWT validation)
    // If authentication is successful, set the authenticated state to true
    setAuthenticated(true);

  };

  const PrivateRoute = ({ element, ...props }) => {
    return authenticated ? element : <Navigate to="/login" />;
  };

  const addToCart = product => {
    const existingItem = cart.find(item => item.product.id === product.id);
    if (existingItem) {
      setCart(prevCart =>
        prevCart.map(item =>
          item.product.id === product.id
            ? { ...item, quantity: item.quantity + 1 }
            : item,
        )
      );
    } else {
      setCart(prevCart => [...prevCart, { product: product, quantity: 1 }]);
    }

    const totalQuantity = cart.reduce((total, item) => {
      if (item.product.id === product.id) {
        total += item.quantity;
      }
      return total;
    }, 1);

    const data = {
      productId: product.id,
      quantity: totalQuantity
    }

    fetch('http://localhost:8080/v1/carts', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        'x-online-store-key': token
      },
      body: JSON.stringify(data)
    })
    .then(response => response.json())
    .then(data => {
      // Handle the response from the backend if needed
      console.log(data);

    })
    .catch(error => {
      // Handle errors
      console.error('Error:', error);
    });

    console.log('updated cart data: ',data)

  };

  const calculateTotal = () => {
    return cart.reduce((total, item) => total + item.product.price * item.quantity, 0);
  };

  const handleConfirm = () => {
    // Prepare the data to send to the backend
    const data = {
      cart: cart.map(item => ({
        productId: item.product.id,
        quantity: item.quantity,
      })),
      total: calculateTotal(),
    };
  
    // Make a POST request to the backend API
    fetch('http://localhost:8080/api/purchase', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(data),
  })
    .then(response => response.json())
    .then(data => {
      // Handle the response from the backend if needed
      console.log(data);

      // Update the purchases state with the new purchase data
      setPurchases(prevPurchases => [...prevPurchases, data]);

      // Reset the cart after successful purchase
      setCart([]);
    })
    .catch(error => {
      // Handle errors
      console.error('Error:', error);
    });
  };
  

  return (
    <Router>
      <div className="App">
        <nav>
          <Link to="/">Home</Link>
          <Link to="/cart">Cart</Link>
          <Link to="/history">History</Link>
        </nav>
        <Routes>
          <Route path="/login" element={<Login onLogin={handleLogin} />} />
          <Route
            path="/"
            element={
              <PrivateRoute
                element={
                  <React.Fragment>
                    <ProductList addToCart={addToCart} />
                    <Cart cart={cart} />
                    <Summary cart={cart} total={calculateTotal()} onConfirm={handleConfirm} />
                    <History purchases={purchases} />
                  </React.Fragment>
                }
              />
            }
          />
        </Routes>
      </div>
    </Router>
  );
};

export default App;
