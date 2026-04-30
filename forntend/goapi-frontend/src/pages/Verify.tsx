import { useEffect, useRef, useState } from 'react';
import { useSearchParams, Link } from 'react-router';
import api from '../services/api';

export default function Verify() {
  const [searchParams] = useSearchParams();
  const [status, setStatus] = useState<'loading' | 'success' | 'error'>('loading');
  const [message, setMessage] = useState('');
  const called = useRef(false);

  useEffect(() => {
    if (called.current) return;
    called.current = true;

    const token = searchParams.get('token');

    if (!token) {
      setStatus('error');
      setMessage('No verification token found in the link.');
      return;
    }

    api
      .get(`/verify?token=${token}`)
      .then((res) => {
        setStatus('success');
        setMessage(res.data.message);
      })
      .catch((err) => {
        setStatus('error');
        setMessage(err.response?.data?.error || 'Verification failed. The link may be invalid or expired.');
      });
  }, [searchParams]);

  return (
    <div className="min-h-screen bg-gray-100 flex items-center justify-center">
      <div className="bg-white p-8 rounded-lg shadow-md w-full max-w-md text-center">
        {status === 'loading' && (
          <>
            <div className="text-4xl mb-4">⏳</div>
            <h1 className="text-xl font-bold text-gray-700">Verifying your email...</h1>
          </>
        )}

        {status === 'success' && (
          <>
            <div className="text-4xl mb-4">✅</div>
            <h1 className="text-2xl font-bold text-green-600 mb-2">Email Verified!</h1>
            <p className="text-gray-500 mb-6">{message}</p>
            <Link
              to="/login"
              className="bg-blue-600 text-white px-6 py-2 rounded-lg hover:bg-blue-700 transition"
            >
              Go to Login
            </Link>
          </>
        )}

        {status === 'error' && (
          <>
            <div className="text-4xl mb-4">❌</div>
            <h1 className="text-2xl font-bold text-red-600 mb-2">Verification Failed</h1>
            <p className="text-gray-500 mb-6">{message}</p>
            <Link
              to="/register"
              className="bg-blue-600 text-white px-6 py-2 rounded-lg hover:bg-blue-700 transition"
            >
              Back to Register
            </Link>
          </>
        )}
      </div>
    </div>
  );
}
