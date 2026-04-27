import { createContext, useContext, useState, type ReactNode } from "react"

interface AuthContextType{
    token: string | null
    login: (token: string, refreshToken: string) => void
    logout: () => void
    isAuthenticated: boolean
}

const AuthContext = createContext <AuthContextType | null >(null)

export function AuthProvider ({children} : {children: ReactNode}){
    const [token, setToken] = useState <string | null>(
        localStorage.getItem(`access_token`)
    )


    const login = (accesToken :string, refreshToken: string) =>{
        localStorage.setItem("access_token", accesToken)
        localStorage.setItem("refresh_token", refreshToken)

        setToken(accesToken)
    }

    const logout = () => {
        localStorage.removeItem(`access_token`)

        localStorage.removeItem(`refresh_token`)

        setToken(null)
    }


    return (
        <AuthContext.Provider value={{
            token, 
            login,
            logout,
            isAuthenticated: !!token
        }} >

            {children}

        </AuthContext.Provider>
    )



}


export function useAuth(){
    const context = useContext(AuthContext)
    if (!context) throw new Error("useAuth must be used within AuthProvider")
    return context
}