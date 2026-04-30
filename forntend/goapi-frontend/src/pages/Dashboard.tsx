import { useEffect, useState } from 'react';
import { useNavigate } from 'react-router';
import { useAuth } from '../context/AuthContext';
import { getUsers, logout, deleteUser } from '../services/api';

interface User {
  ID: number;
  name: string;
  email: string;
  age: number;
}

export default function Dashboard() {
  const { isAuthenticated, logout: authLogout } = useAuth();
  const navigate = useNavigate();

  const [users, setUsers] = useState<User[]>([]);
  const [loading, setLoading] = useState(false);
  const [deletingId, setDeletingId] = useState<number | null>(null);

  useEffect(() => {
    if (!isAuthenticated) {
      navigate('/login');
      return;
    }

    // 👇 Define and call inside useEffect
    const loadUsers = async () => {
      try {
        const res = await getUsers();
        setUsers(res.data.data);
      } catch (err) {
        console.error(err);
      } finally {
        setLoading(false);
      }
    };

    loadUsers();
  }, [isAuthenticated, navigate]);

  const handleDelete = async (id: number) => {
    setDeletingId(id);
    try {
      await deleteUser(id);
      setUsers((prev) => prev.filter((u) => u.ID !== id));
    } catch (err) {
      console.error(err);
    } finally {
      setDeletingId(null);
    }
  };

  const handleLogout = async () => {
    try {
      await logout();
    } catch (err) {
      console.error(err);
    } finally {
      authLogout();
      navigate('/login');
    }
  };

  return (
    <div className="min-h-screen bg-gray-100">
      <nav className="bg-white shadow-sm px-6 py-4 flex justify-between items-center">
        <h1 className="text-xl font-bold text-blue-600 ">GoAPI Dashboard</h1>

        <button onClick={handleLogout}>Logout</button>
      </nav>

      <div className="max-w-4xl mx-auto p-6">
        <h2 className="text-2xl font-bold text-gray-800 mb-6">All Users</h2>
        {loading ? (
          <div className="text-center text-gray-500">Loading...</div>
        ) : (
          <div className="grid gap-4">
            {users.map((user) => (
              <div
                key={user.ID}
                className="bg-white p-4 rounded-lg shadow-sm flex items-center justify-between"
              >
                <div>
                  <p className="font-medium text-gray-800">{user.name}</p>
                  <p className="text-sm text-gray-500">{user.email}</p>
                </div>

                <div className="flex items-center gap-3">
                  <span className="bg-blue-100 text-blue-600 px-3 py-1 rounded-full text-sm">
                    Age: {user.age}
                  </span>
                  <button
                    onClick={() => handleDelete(user.ID)}
                    disabled={deletingId === user.ID}
                    className="bg-red-500 text-white px-3 py-1 rounded-lg text-sm hover:bg-red-600 transition disabled:opacity-50"
                  >
                    {deletingId === user.ID ? 'Deleting...' : 'Delete'}
                  </button>
                </div>
              </div>
            ))}
          </div>
        )}
      </div>
    </div>
  );
}
