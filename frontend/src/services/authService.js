import axios from 'axios';

const BASE_URL = 'http://localhost:8080';

// Register a new user
export const registerUser = async (username, password, role) => {
  try {
    const response = await axios.post(`${BASE_URL}/register`, {
      username,
      password,
      role,
    });
    return response.data;
  } catch (err) {
    throw new Error(err.response?.data?.error || 'Registration failed');
  }
};

// Login user and return token
export const loginUser = async (username, password) => {
  try {
    const response = await axios.post(`${BASE_URL}/login`, {
      username,
      password,
    });
    return response.data;
  } catch (err) {
    throw new Error(err.response?.data?.error || 'Login failed');
  }
};

// Get token from localStorage
export const getToken = () => {
  return localStorage.getItem('token');
};

// Logout — clear token
export const logoutUser = () => {
  localStorage.removeItem('token');
};

// Axios instance with token attached for protected routes
export const authAxios = axios.create({
  baseURL: BASE_URL,
});

authAxios.interceptors.request.use((config) => {
  const token = getToken();
  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }
  return config;
});