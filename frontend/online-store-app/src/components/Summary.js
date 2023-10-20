import React from 'react';

const Summary = ({ cart, total, handleConfirm }) => {
  return (
    <div className="summary">
      <h2>Order Summary</h2>
      <div className="cart-summary">
        {cart.map(item => (
          <div key={item.product.id} className="cart-item">
            <p>{item.product.name}</p>
            <p>Quantity: {item.quantity}</p>
          </div>
        ))}
      </div>
      <h3>Total: ${total}</h3>
      <button onClick={handleConfirm}>Confirm Purchase</button>
    </div>
  );
};

export default Summary;
