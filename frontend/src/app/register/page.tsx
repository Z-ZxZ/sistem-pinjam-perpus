'use client';

import React, { useState } from 'react';
import { useRouter } from 'next/navigation';
import { Button } from '@/components/ui/Button';
import { api } from '@/lib/api';
import { Book, Mail, Lock, User, ArrowLeft } from 'lucide-react';
import Link from 'next/link';

export default function RegisterPage() {
  const router = useRouter();
  const [name, setName] = useState('');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState('');
  const [isLoading, setIsLoading] = useState(false);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setIsLoading(true);
    setError('');
    
    try {
      await api.post('/auth/register', { name, email, password });
      router.push('/login?registered=true');
    } catch (err: any) {
      setError(err.message);
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
            <div className="w-16 h-16 bg-[#14B8A6] rounded-2xl flex items-center justify-center">
              <Book className="text-white w-8 h-8" />
            </div>
          </div>
          
          <h2 className="text-3xl font-bold text-center text-[#1E293B] mb-2">Buat Akun</h2>
          <p className="text-[#64748B] text-center mb-10">Daftar untuk mulai meminjam buku</p>

          {error && (
            <div className="bg-red-50 text-red-600 p-4 rounded-xl text-sm mb-6 border border-red-100">
              {error}
            </div>
          )}

          <form onSubmit={handleSubmit} className="space-y-5">
            <div className="space-y-2">
              <label className="text-sm font-semibold text-[#1E293B] block">Nama Lengkap</label>
              <div className="relative">
                <div className="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                  <User className="h-5 w-5 text-[#94A3B8]" />
                </div>
                <input
                  type="text"
                  required
                  value={name}
                  onChange={(e) => setName(e.target.value)}
                  className="w-full pl-10 pr-4 py-3 bg-[#F8FAFC] border border-[#E2E8F0] rounded-xl focus:ring-2 focus:ring-[#4338CA]/20 focus:border-[#4338CA] transition-all"
                  placeholder="Nama lengkap Anda"
                />
              </div>
            </div>

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
                  placeholder="Minimal 8 karakter"
                />
              </div>
            </div>

            <Button type="submit" variant="accent" isLoading={isLoading} className="w-full py-4 text-lg mt-4">
              Daftar
            </Button>
          </form>

          <p className="text-center mt-8 text-[#64748B]">
            Sudah punya akun?{' '}
            <Link href="/login" className="text-[#14B8A6] font-semibold hover:underline">
              Masuk
            </Link>
          </p>
        </div>
      </div>
    </div>
  );
}
