
import { useState } from "react"
import { register } from "../services/api"
import { Link } from "react-router"
//import { useNavigate } from "react-router"


export default function Register() {
    //const negative = useNavigate()
    const [form, setForm] = useState({
        name: "",
        email: "",
        password: "",
        age: "",
    })

    const [error, setError] = useState("")
    const [success, setSuccess] = useState ("")
    const [loading, setLoading] = useState(false)

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault()
        setLoading(true)
        setError("")

        try{
            await register({
                name: form.name,
                email: form.email,
                password: form.password,
                age: parseInt(form.age)
            })

            setSuccess("Registration seccessful! Check you email to verify you account.")
        }catch(err: any){

            setError(err.response?.data?.error || "somthing went wrong")
        }finally{
            setLoading(false)
        }
    }

  return (
    
    <div className='min-h-screen bg-gray-100 flex items-center justify-center'>
        <div className='bg-white p-8 rounded-lg shadow-lg w-full max-w-md'>
            <h1 className='text-3xl font-bold text-center text-b'>
                Create Account
            </h1>

            <p className='text-center text-gray-500 mb-6'>Join GoApi today</p>

            {error && (
                <div className="bg-red-50 text-red-600 p-3 rounded-lg mb-4 text-sm">
                    {error} 
                </div>
            )}

            {success && (
                <div className="bg-green-50 text-green-600 p-3 rounded-lg mb-4 text-sm">
                    {success}
                </div>
            )}

            <form onSubmit={handleSubmit} className="space-y-4">
                <div>
                    <label className="block text-sm front-medum text-gray-700 mb-1">
                        Full Name
                    </label>

                    <input 
                        type="text"
                        placeholder="Dachi"
                        value={form.name} 
                        onChange={(e) => setForm({...form, name: e.target.value})} className="w-full border border-gray-300 rounded-lg px-4 py-2 focus:outline-none focus:ring-2 focus:ring-blue-500" 
                    />
                </div>

                <div>
                    <label className="block text-sm font-medium text-gray-700 mb-1">
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
                    placeholder="••••••••"
                    value={form.password}
                    onChange={(e) => setForm({ ...form, password: e.target.value })}
                    className="w-full border border-gray-300 rounded-lg px-4 py-2 focus:outline-none focus:ring-2 focus:ring-blue-500"
                    />
                </div>

                <div>
                    <label className="block text-sm font-medium text-gray-700 mb-1">
                    Age
                    </label>
                    <input
                    type="number"
                    placeholder="25"
                    value={form.age}
                    onChange={(e) => setForm({ ...form, age: e.target.value })}
                    className="w-full border border-gray-300 rounded-lg px-4 py-2 focus:outline-none focus:ring-2 focus:ring-blue-500"
                    />
                </div>

                <button 
                    type="submit"
                    disabled={loading} 
                    className="w-full bg-blue-600 text-white py-2 rounded-ld font-medium hover:bg-blue-700 transition disabled:opacity-50"
                >
                    {loading ? "Creating account..." : "Regiter"}
                </button>
            </form>

            <p className="text-center text-gray-500 text-sm mt-4">
                Already have an account? 
                <Link to="/login" className="text-blue-600  hover:underline"
                >
                    Login
                </Link>
            </p>
        </div>
    </div>
  )
}
