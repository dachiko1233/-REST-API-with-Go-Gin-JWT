import { useState } from "react"
import { useAuth } from "../context/AuthContext"
import { data, Link, useNavigate } from "react-router"
import { login } from "../services/api"


export default function Login() {
    const negative = useNavigate()
    const [loading, setLoading] = useState(false)
    const {login: authLogin} = useAuth()
    const [form, setForm] = useState({email: "", password: ""})
    const [error, setError] =useState("")

    const handleSubmit = async(e: React.FormEvent) => {
        e.preventDefault()
        setLoading(true)
        setError('')

        try{
            const res = await login(form)
            authLogin(res.data.token, res.data.refresh_token)
            negative("/dashboard")
        }catch(err:any){
            setError(err.response?.data?.error || "somthing went wrong")
        }finally{
            setLoading(false)
        }
    }

  return (
    <div  className="min-h-screen bg-gray-100 flex items-center justify-center">
        <div className="bg-white p-8 rounded-lg shadow-md w-full max-w-md">

            <h1 className="text-center text-3xl font-bold text-blue-600 mb-2 ">Welcome Back!</h1>

            <p className="text-center text-gray-500 mb-6">Logi to your account</p>

            {error && (
                <div className="bg-red-50 text-red-600 p-3 rounded-lg mb-4 text-sm">
                    {error}
                </div>
            )}

            <form onSubmit={handleSubmit} className="space-y-4">
                <div>
                    <label className="block text-sm font-medium text-gray-700 mb-1 ">
                        Email
                    </label>
                   
                    <input
                        type="email"
                        placeholder="dachi@example.com"
                        value={form.email}
                        onChange={(e) => setForm({ ...form, email: e.target.value })}
                        className="w-full border border-gray-300 rounded-lg px-4 py-2 focus:outline-none focus:ring-2 focus:ring-blue-500"
                    />


                </div>

                <div>
                    <label className="block text-sm font-medium text-gray-700 mb-1">
                        Password
                    </label>
                    <input
                     type="password"
                     value={form.password}
                     onChange={(e) => setForm({...form, password: e.target.value})}
                     className="w-full border border-gray-300 rounded-lg px-4 py-2 focus:outline-none focus:ring-2 focus:ring-blue-500" 
                     placeholder="••••••••"

                     />
                </div>


                <button
                    type="submit"
                    disabled={loading} 
                    className="w-full bg-blue-600 text-white py-2 rounded-lg font-medium hover:bg-blue-700 transition disabled:opacity-50"
                >
                    {loading ? "Loggin in..." : "Login"}
                </button>

            </form>
            <p className="text-center text-gray-500 text-sm mt-4">
                Don't have an account?{' '}
                <Link to="/register" className="text-blue-600           hover:underline"
                >
                    Register
                </Link>
            </p>
        </div>
    </div>
  )
}
