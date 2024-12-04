import axios from 'axios';
const baseURL = 'https://wombat-production-e2c6.up.railway.app/api';

const api = axios.create({
    baseURL: baseURL,
});

export default api;