import React from 'react';
import { Line } from 'react-chartjs-2';
import { Chart as ChartJS, CategoryScale, LinearScale, PointElement, LineElement, Title, Tooltip, Legend } from 'chart.js';

ChartJS.register(CategoryScale, LinearScale, PointElement, LineElement, Title, Tooltip, Legend);

const LightsChart = ({ data }) => {
  const chartData = {
    labels: data.map(item => item.event_time),
    datasets: [{
      label: 'Lights',
      data: data.map(item => item.value_int),
      fill: false,
      borderColor: 'rgb(75, 192, 192)',
      tension: 0.1,
    }],
  };

  const options = {
    scales: {
      y: {
        type: 'category', // Defina o tipo como categoria
        labels: ['Ligado', 'Desligado'], // Defina as categorias explicitamente
        reverse: true,
      }
    }
  };

  return (
    <div>
      <h2>Lights</h2>
      <Line data={chartData} options={options} />
    </div>
  );
};

export default LightsChart;