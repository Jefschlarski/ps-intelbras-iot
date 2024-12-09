"use client";
import React from 'react';
import { useParams } from 'next/navigation';
import DeviceChart from '@/app/components/DeviceChart/DeviceChart';

const DeviceChartsPage = () => {
  const { deviceId } = useParams(); // Obtém o deviceId da URL

  if (!deviceId) {
    return <p>Carregando...</p>;
  }

  return (
    <div>
      <DeviceChart deviceId={parseInt(deviceId)} />
    </div>
  );
};

export default DeviceChartsPage;
