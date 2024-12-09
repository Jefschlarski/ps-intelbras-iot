import React from 'react';
import { Line } from 'react-chartjs-2';
import { Chart as ChartJS, CategoryScale, LinearScale, PointElement, LineElement, Title, Tooltip, Legend } from 'chart.js';

ChartJS.register(CategoryScale, LinearScale, PointElement, LineElement, Title, Tooltip, Legend);

const MileageChart = ({ data }) => {
  const chartData = {
    labels: data.map(item => item.event_time),
    datasets: [{
      label: 'Mileage',
      data: data.map(item => item.value_int), // Dados tipo int
      fill: false,
      borderColor: 'rgb(75, 192, 192)',
      tension: 0.5, 
      borderJoinStyle: 'round', 
      pointRadius: 1, 
    }],
  };

  return (
    <div>
      <h2>Mileage</h2>
      <Line data={chartData} />
    </div>
  );
};

export default MileageChart;
