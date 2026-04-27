import axios from "axios";

const api = axios.create({
    baseURL:" http://localhost:8080/api/v1",
    headers:{
        'Content-Type': 'application/json',
    },
})

// Automatically add token to every request
api.interceptors.request.use((config) => {
    const token = localStorage.getItem("access_token")
    if (token){
        config.headers.Authorization = `Bearer ${token}`
    }

    return config
})


// Auth endpoints

export const register = (data:{
    name: string
    email:string
    password: string
    age: number
}) => api.post("/register", data)

export const login = (data :{
    email:string
    password: string

}) => api.post('/login', data)


export const logout = () => api.post('/logout')

export const  getUsers = () => api.get('/users')

export default api