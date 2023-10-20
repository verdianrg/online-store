import React, { useState, useEffect } from 'react';

const ProductList = ({ addToCart }) => {
  const [products, setProducts] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  const token = localStorage.getItem('token')

  useEffect(() => {
    fetch('http://localhost:8080/v1/products', {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
        'x-online-store-key': token
      }
    })
      .then(response => response.json())
      .then(data => {
        console.log('data: ', data)
        setProducts(data.data);
        setLoading(false); // Set loading to false after data is fetched successfully
      })
      .catch(error => {
        setError(error); // Set error state in case of fetch failure
        setLoading(false); // Set loading to false in case of fetch failure
      });
  }, []); // Empty dependency array ensures the effect runs once after the initial render

  if (loading) {
    return <div>Loading...</div>;
  }

  if (error) {
    return <div>Error: {error.message}</div>;
  }

  return (
    <div>
      <h2>Product List</h2>
      <ul>
        {products.map(product => (
          <li key={product.id}>
            {product.name} - {product.price}{' '}
            <button onClick={() => addToCart(product)}>Add to Cart</button>
          </li>
        ))}
      </ul>
    </div>
  );
};

export default ProductList;
