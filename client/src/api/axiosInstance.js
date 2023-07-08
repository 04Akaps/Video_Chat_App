import axios from 'axios';
import { Cookies } from 'react-cookie';

const axiosInstance = axios.create({
    baseURL: `http://localhost:8000`,
});

axiosInstance.interceptors.request.use((config) => {
    const cookies = new Cookies();
    const oauthCookie = cookies.get('oauth');
    if (oauthCookie) {
        config.headers.Authorization = `Bearer ${oauthCookie}`;
    }
    return config;
})

export default axiosInstance