import React from 'react';

const History = ({ purchases }) => {
  return (
    <div className="history">
      <h2>Purchase History</h2>
      {purchases.map((purchase, index) => (
        <div key={index} className="purchase-item">
          <p>{purchase.date}</p>
          <p>Status: {purchase.status}</p>
        </div>
      ))}
    </div>
  );
};

export default History;
