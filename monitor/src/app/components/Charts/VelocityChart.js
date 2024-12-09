import React from 'react';
import { Line } from 'react-chartjs-2';
import { Chart as ChartJS, CategoryScale, LinearScale, PointElement, LineElement, Title, Tooltip, Legend } from 'chart.js';

ChartJS.register(CategoryScale, LinearScale, PointElement, LineElement, Title, Tooltip, Legend);

const VelocityChart = ({ data }) => {
  const chartData = {
    labels: data.map(item => item.event_time),
    datasets: [{
      label: 'Velocity',
      data: data.map(item => item.value_float), // Dados tipo float
      fill: false,
      borderColor: 'rgb(75, 192, 192)',
      tension: 0.5, 
      borderJoinStyle: 'round', 
      pointRadius: 1, 
    }],
  };

  return (
    <div>
      <h2>Velocity</h2>
      <Line data={chartData} />
    </div>
  );
};

export default VelocityChart;
