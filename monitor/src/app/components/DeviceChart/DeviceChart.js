import React, { useEffect, useState } from 'react';
import { useRouter } from 'next/navigation';
import VelocityChart from '../Charts/VelocityChart';
import MileageChart from '../Charts/MileageChart';
import LightsChart from '../Charts/LightsChart';
import GpsChart from '../Charts/GpsChart';
import ErrorChart from '../Charts/ErrorChart';
import RpmChart from '../Charts/RpmChart';
import TemperatureChart from '../Charts/TemperatureChart';
import FuelChart from '../Charts/FuelChart';
import styles from './DeviceChart.module.css';

const DeviceChart = ({ deviceId }) => {
  const [chartsData, setChartsData] = useState({});
  const router = useRouter();

  useEffect(() => {
    const fetchChartsData = async () => {
      try {
        const responses = await Promise.all([
          fetch(`http://localhost:5555/api/v1/telemetries/devices/${deviceId}?type=1&limit=100`), // Velocity
          fetch(`http://localhost:5555/api/v1/telemetries/devices/${deviceId}?type=2&limit=100`), // RPM
          fetch(`http://localhost:5555/api/v1/telemetries/devices/${deviceId}?type=3&limit=100`), // Temp
          fetch(`http://localhost:5555/api/v1/telemetries/devices/${deviceId}?type=4&limit=100`), // Fuel
          fetch(`http://localhost:5555/api/v1/telemetries/devices/${deviceId}?type=5&limit=100`), // Mileage
          fetch(`http://localhost:5555/api/v1/telemetries/devices/${deviceId}?type=6&limit=100`), // GPS
          fetch(`http://localhost:5555/api/v1/telemetries/devices/${deviceId}?type=7&limit=100`), // Lights
          fetch(`http://localhost:5555/api/v1/telemetries/devices/${deviceId}?type=8&limit=100`), // Error
        ]);

        const data = await Promise.all(responses.map(async response => {
          if (!response.ok) {
            throw new Error(`Erro ${response.status}: ${response.statusText}`);
          }

          const contentType = response.headers.get('Content-Type');
          if (contentType && contentType.includes('application/json')) {
            const text = await response.text();
            return text ? JSON.parse(text) : []; // Verifica se o corpo não está vazio
          } else {
            throw new Error('Resposta não é JSON');
          }
        }));

        setChartsData({
          velocity: data[0],
          rpm: data[1],
          temp: data[2],
          fuel: data[3],
          mileage: data[4],
          gps: data[5],
          lights: data[6],
          error: data[7],
        });
      } catch (error) {
        console.error('Erro ao carregar gráficos:', error);
      }
    };

    fetchChartsData();
  }, [deviceId]);

  const handleBack = () => {
    router.push('/');
  };

  return (
    <div className={styles.container}>
      <h1 className={styles.title}>Gráficos do Dispositivo {deviceId}</h1>
      <button className={styles.backButton} onClick={handleBack}>
        Voltar
      </button>
      <div className={styles.grid}>
        {chartsData.velocity && (
          <div className={styles.chartItem}>
            <VelocityChart data={chartsData.velocity} />
          </div>
        )}
        {chartsData.rpm && (
          <div className={styles.chartItem}>
            <RpmChart data={chartsData.rpm} />
          </div>
        )}
        {chartsData.temp && (
          <div className={styles.chartItem}>
            <TemperatureChart data={chartsData.temp} />
          </div>
        )}
        {chartsData.fuel && (
          <div className={styles.chartItem}>
            <FuelChart data={chartsData.fuel} />
          </div>
        )}
        {chartsData.mileage && (
          <div className={styles.chartItem}>
            <MileageChart data={chartsData.mileage} />
          </div>
        )}
        {chartsData.gps && (
          <div className={styles.chartItem}>
            <GpsChart data={chartsData.gps} />
          </div>
        )}
        {chartsData.lights && (
          <div className={styles.chartItem}>
            <LightsChart data={chartsData.lights} />
          </div>
        )}
        {chartsData.error && (
          <div className={styles.chartItem}>
            <ErrorChart data={chartsData.error} />
          </div>
        )}
      </div>
    </div>
  );
};

export default DeviceChart;
