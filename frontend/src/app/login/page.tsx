'use client';

import React, { useState, useEffect } from 'react';
import { useRouter } from 'next/navigation';
import { Button } from '@/components/ui/Button';
import { api } from '@/lib/api';
import { Book, Mail, Lock, ArrowLeft } from 'lucide-react';
import Link from 'next/link';

import { useAuth } from '@/context/AuthContext';

export default function LoginPage() {
  const router = useRouter();
  const { login: authLogin, isLoggedIn, user, isRestored } = useAuth();
  
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState('');
  const [isLoading, setIsLoading] = useState(false);

  useEffect(() => {
    if (isRestored && isLoggedIn && user) {
      console.log('[Login] Session detected, redirecting...', user.role);
      if (user.role === 'admin') {
        router.push('/admin');
      } else {
        router.push('/dashboard');
      }
    }
  }, [isRestored, isLoggedIn, user, router]);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setIsLoading(true);
    setError('');
    
    try {
      const res = await api.post('/auth/login', { email, password });

      console.log('[Login] API Response:', res);

      const userData = res.user;

      authLogin(res.token, userData);
      // Bakal di handle sama useEffect ntar
    } catch (err: unknown) {
      if (err instanceof Error) {
        setError(err.message);
      } else {
        setError('Login gagal. Periksa kembali email dan password Anda.');
      }
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <div className="min-h-screen bg-[#F8FAFC] flex items-center justify-center px-4 py-12">
      <div className="max-w-md w-full">
        <Link href="/" className="inline-flex items-center gap-2 text-[#64748B] hover:text-[#4338CA] transition-colors mb-8 group">
          <ArrowLeft size={18} className="group-hover:-translate-x-1 transition-transform" />
          <span>Kembali ke Beranda</span>
        </Link>
        
        <div className="bg-white rounded-3xl p-10 shadow-xl border border-[#E2E8F0]">
          <div className="flex justify-center mb-8">
            <div className="w-16 h-16 bg-[#4338CA] rounded-2xl flex items-center justify-center">
              <Book className="text-white w-8 h-8" />
            </div>
          </div>
          
          <h2 className="text-3xl font-bold text-center text-[#1E293B] mb-2">Selamat Datang</h2>
          <p className="text-[#64748B] text-center mb-10">Masuk untuk mengakses layanan perpustakaan</p>

          {error && (
            <div className="bg-red-50 text-red-600 p-4 rounded-xl text-sm mb-6 border border-red-100">
              {error}
            </div>
          )}

          <form onSubmit={handleSubmit} className="space-y-6">
            <div className="space-y-2">
              <label className="text-sm font-semibold text-[#1E293B] block">Email</label>
              <div className="relative">
                <div className="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                  <Mail className="h-5 w-5 text-[#94A3B8]" />
                </div>
                <input
                  type="email"
                  required
                  value={email}
                  onChange={(e) => setEmail(e.target.value)}
                  className="w-full pl-10 pr-4 py-3 bg-[#F8FAFC] border border-[#E2E8F0] rounded-xl focus:ring-2 focus:ring-[#4338CA]/20 focus:border-[#4338CA] transition-all"
                  placeholder="nama@email.com"
                />
              </div>
            </div>

            <div className="space-y-2">
              <label className="text-sm font-semibold text-[#1E293B] block">Kata Sandi</label>
              <div className="relative">
                <div className="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                  <Lock className="h-5 w-5 text-[#94A3B8]" />
                </div>
                <input
                  type="password"
                  required
                  value={password}
                  onChange={(e) => setPassword(e.target.value)}
                  className="w-full pl-10 pr-4 py-3 bg-[#F8FAFC] border border-[#E2E8F0] rounded-xl focus:ring-2 focus:ring-[#4338CA]/20 focus:border-[#4338CA] transition-all"
                  placeholder="••••••••"
                />
              </div>
            </div>

            <Button type="submit" isLoading={isLoading} className="w-full py-4 text-lg">
              Masuk
            </Button>
          </form>

          <p className="text-center mt-8 text-[#64748B]">
            Belum punya akun?{' '}
            <Link href="/register" className="text-[#4338CA] font-semibold hover:underline">
              Daftar Sekarang
            </Link>
          </p>
        </div>
      </div>
    </div>
  );
}
