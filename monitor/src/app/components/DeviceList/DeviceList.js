"use client";
import React, { useEffect, useState } from 'react';
import { useRouter } from 'next/navigation';
import styles from './DeviceList.module.css';

const DeviceList = () => {
  const [devices, setDevices] = useState([]);
  const [telemetryCount, setTelemetryCount] = useState(0);
  const [isClient, setIsClient] = useState(false);
  const router = useRouter();

  useEffect(() => {
    setIsClient(true);
  }, []);

  useEffect(() => {
    const fetchDevices = async () => {
      try {
        const response = await fetch('http://localhost:5555/api/v1/telemetries/devices');
        const data = await response.json();
        setDevices(data);
      } catch (error) {
        console.error('Erro ao carregar dispositivos:', error);
      }
    };

    fetchDevices();
  }, []);

  useEffect(() => {
    const fetchTelemetryCount = async () => {
      try {
        const response = await fetch('http://localhost:5555/api/v1/telemetries/length');
        const count = await response.json();
        setTelemetryCount(count);
      } catch (error) {
        console.error('Erro ao carregar contagem de telemetrias:', error);
      }
    };

    // Atualiza a cada 1 segundo
    const interval = setInterval(fetchTelemetryCount, 1000);

    // Busca inicial
    fetchTelemetryCount();

    // Limpa o intervalo ao desmontar
    return () => clearInterval(interval);
  }, []);

  const handleViewCharts = (deviceId) => {
    router.push(`/device-charts/${deviceId}`);
  };

  if (!isClient) {
    return (
      <div className={styles.loading}>
        <p>Carregando...</p>
      </div>
    );
  }

  return (
    <div className={styles.container}>
      <p className={styles.telemetryCount}>
        <strong>Total de Telemetrias:</strong> {telemetryCount}
      </p>
      <h1 className={styles.title}>Dispositivos Disponíveis</h1>
      <ul className={styles.deviceList}>
        {devices.map((device) => (
          <li key={device.device_id} className={styles.deviceItem}>
            <div className={styles.deviceInfo}>
              <p><strong>Device ID:</strong> {device.device_id}</p>
              <p><strong>Última Telemetria:</strong> {device.last_telemetry_time}</p>
              <p><strong>Último Tipo de Evento:</strong> {device.last_telemetry_type}</p>
            </div>
            <button className={styles.viewButton} onClick={() => handleViewCharts(device.device_id)}>
              Ver Gráficos
            </button>
          </li>
        ))}
      </ul>
    </div>
  );
};

export default DeviceList;
