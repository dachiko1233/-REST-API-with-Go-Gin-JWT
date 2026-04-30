import { BrowserRouter, Route, Routes } from "react-router";
import { AuthProvider } from "./context/AuthContext";

import Register from "./pages/Register";
import Login from "./pages/Login";
import Dashboard from "./pages/Dashboard";
import Verify from "./pages/Verify";

export function App(){
  return(
    <AuthProvider>
      <BrowserRouter>
      <Routes>
        
        <Route path="/login" element={<Login/>} />
        <Route path="/register" element={<Register/>} />
        <Route path="/dashboard" element={<Dashboard/>} />
        <Route path="/verify" element={<Verify/>} />

      </Routes>
      </BrowserRouter>
    </AuthProvider>
  )
}