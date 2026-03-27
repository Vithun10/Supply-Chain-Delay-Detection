import axios from 'axios';

const api = axios.create({
    baseURL: 'http://localhost:8080',
    timeout: 10000,
    headers: {
        'Content-Type': 'application/json',
    }
});

// Automatically attach JWT token to every request
api.interceptors.request.use((config) => {
    const token = localStorage.getItem('token');
    if (token) {
        config.headers.Authorization = `Bearer ${token}`;
    }
    return config;
});

// Shipment APIs
export const fetchShipments = (params) => api.get('/shipments', { params }).then(res => res.data);
export const fetchShipmentById = (id) => api.get(`/shipments/${id}`).then(res => res.data);

// Delay APIs
export const fetchDelays = (params) => api.get('/delays', { params }).then(res => res.data);
export const fetchHighRiskDelays = (params) => api.get('/delays/high-risk', { params }).then(res => res.data);

// Analytics APIs
export const fetchDelayRate = () => api.get('/analytics/delay-rate').then(res => res.data);
export const fetchTopDelayedRoutes = () => api.get('/analytics/top-delayed-routes').then(res => res.data);
export const fetchCarrierPerformance = () => api.get('/analytics/carrier-performance').then(res => res.data);
export const fetchAvgDeliveryTime = () => api.get('/analytics/avg-delivery-time').then(res => res.data);

// Tracking API
export const trackShipment = (id) => api.get(`/track/${id}`).then(res => res.data);

export default api;