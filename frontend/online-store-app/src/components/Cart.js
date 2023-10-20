import React from 'react';

const Cart = ({ cart }) => {
  console.log(cart);
  return (
    <div className="cart">
      <h2>Shopping Cart</h2>
      {cart.map(item => (
        <div key={item.product.id} className="cart-item">
          <p>{item.product.name}</p>
          <p>Quantity: {item.quantity}</p>
        </div>
      ))}
    </div>
  );
};

export default Cart;
