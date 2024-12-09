import React from 'react';
import { Bar } from 'react-chartjs-2';
import { Chart as ChartJS, CategoryScale, LinearScale, BarElement, Title, Tooltip, Legend } from 'chart.js';
import moment from 'moment';

ChartJS.register(CategoryScale, LinearScale, BarElement, Title, Tooltip, Legend);

const GpsChart = ({ data }) => {
  // Função para agrupar por minuto
  const groupByMinute = (data) => {
    const groupedData = {};

    data.forEach(item => {
      const minute = moment(item.event_time).startOf('minute').format(); // Arredonda para o início do minuto
      if (groupedData[minute]) {
        groupedData[minute]++;
      } else {
        groupedData[minute] = 1;
      }
    });

    return groupedData;
  };

  const groupedData = groupByMinute(data);

  // Preparando os dados para o gráfico
  const chartData = {
    labels: Object.keys(groupedData), // Minutos
    datasets: [{
      label: 'GPS',
      data: Object.values(groupedData), // Contagem de locations por minuto
      backgroundColor: 'rgba(75, 192, 192, 0.2)',
      borderColor: 'rgba(75, 192, 192, 1)',
      borderWidth: 1,
    }],
  };

  return (
    <div>
      <h2>GPS</h2>
      <Bar data={chartData} />
    </div>
  );
};

export default GpsChart;
