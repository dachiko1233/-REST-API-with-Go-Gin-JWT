import { BrowserRouter, Route, Routes } from "react-router-dom";
import { AuthProvider } from "./context/AuthContext";

import Register from "./pages/Register";
import Login from "./pages/Login";

export function App(){
  return(
    <AuthProvider>
      <BrowserRouter>
      <Routes>
        
        <Route path="/login" element={<Login/>}   /> 
        <Route path="/register" element={<Register/>} />
        {/* <Route path="/dashboard" element={<Dashboard/>} />  */}

      </Routes>
      </BrowserRouter>
    </AuthProvider>
  )
}