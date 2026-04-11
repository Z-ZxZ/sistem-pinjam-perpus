'use client';

import React, { useEffect, useState } from 'react';
import { Navbar } from '@/components/layout/Navbar';
import { Button } from '@/components/ui/Button';
import { api } from '@/lib/api';
import { Users, Book as BookIcon, History, TrendingUp, Plus, Trash2, Edit } from 'lucide-react';
import { motion } from 'framer-motion';

interface User {
  id: number;
  name: string;
  email: string;
}

interface Book {
  id: number;
  title: string;
  total?: number;
}

interface Borrow {
  id: number;
  user: User;
  book: Book;
  borrow_date: string;
  status: string;
}

import { useRouter } from 'next/navigation';
import { useAuth } from '@/context/AuthContext';

export default function AdminDashboard() {
  const router = useRouter();
  const { isLoggedIn, isAdmin, isLoading: authLoading } = useAuth();
  
  const [users, setUsers] = useState<User[]>([]);
  const [borrows, setBorrows] = useState<Borrow[]>([]);
  const [booksCount, setBooksCount] = useState(0);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  const fetchData = async () => {
    try {
      const [usersRes, borrowsRes, booksRes] = await Promise.all([
        api.get('/admin/users'),
        api.get('/admin/borrows'),
        api.get('/books?limit=1'), // Just to get total
      ]);
      console.log('[Admin] Fetched users:', usersRes.data.length);
      setUsers(usersRes.data);
      setBorrows(borrowsRes.data);
      setBooksCount(booksRes.data.total);
    } catch (err: unknown) {
      console.error('[Admin] Fetch error:', err);
      if (err instanceof Error) setError(err.message);
      else setError('Gagal memuat data administrasi.');
    } finally {
      setIsLoading(false);
    }
  };

  useEffect(() => {
    if (isRestored) {
      if (!isLoggedIn) {
        console.log('[Admin] Not logged in, redirecting...');
        router.push('/login');
      } else if (!isAdmin) {
        console.log('[Admin] Not an admin, redirecting...');
        router.push('/dashboard');
      } else {
        fetchData();
      }
    }
  }, [isLoggedIn, isAdmin, authLoading, isRestored, router]);

  if (!isRestored || (isLoading && isLoggedIn && isAdmin)) {
    return (
      <div className="min-h-screen bg-[#F8FAFC] flex flex-col items-center justify-center gap-4">
        <div className="w-12 h-12 border-4 border-[#4338CA] border-t-transparent rounded-full animate-spin" />
        <p className="text-[#64748B] font-medium animate-pulse">Menyiapkan Dashboard...</p>
      </div>
    );
  }

  if (error) {
    return (
      <main className="min-h-screen bg-[#F8FAFC]">
        <Navbar />
        <div className="max-w-7xl mx-auto px-4 pt-32 flex flex-col items-center justify-center gap-4">
           <div className="bg-red-50 text-red-600 p-6 rounded-2xl border border-red-100 max-w-md text-center">
             <h2 className="text-xl font-bold mb-2">Terjadi Kesalahan</h2>
             <p>{error}</p>
             <Button className="mt-4" onClick={fetchData}>Coba Lagi</Button>
           </div>
        </div>
      </main>
    );
  }

  if (!isLoggedIn || !isAdmin) return null;

  return (
    <main className="min-h-screen bg-[#F8FAFC]">
      <Navbar />
      
      <div className="max-w-7xl mx-auto px-4 pt-32 pb-20">
        <div className="flex justify-between items-end mb-10">
          <div>
            <h1 className="text-3xl font-bold text-[#1E293B] mb-2">Panel Administrasi</h1>
            <p className="text-[#64748B]">Pantau dan kelola seluruh aktivitas perpustakaan.</p>
          </div>
          <Button variant="accent" className="flex items-center gap-2">
            <Plus size={18} />
            Tambah Buku
          </Button>
        </div>

        {/* Analytics Grid */}
        <div className="grid grid-cols-1 md:grid-cols-4 gap-6 mb-12">
          {[
            { label: 'Total Buku', value: booksCount, icon: BookIcon, color: 'bg-[#4338CA]' },
            { label: 'Total Anggota', value: users.length, icon: Users, color: 'bg-[#14B8A6]' },
            { label: 'Buku Dipinjam', value: borrows.filter(b => b.status === 'borrowed').length, icon: History, color: 'bg-orange-500' },
            { label: 'Tren Bulanan', value: '+12%', icon: TrendingUp, color: 'bg-green-500' },
          ].map((stat, i) => (
            <motion.div 
              key={stat.label}
              initial={{ opacity: 0, scale: 0.95 }}
              animate={{ opacity: 1, scale: 1 }}
              transition={{ delay: i * 0.1 }}
              className="card bg-white"
            >
              <div className="flex items-center gap-4">
                <div className={`w-12 h-12 ${stat.color} rounded-xl flex items-center justify-center text-white shadow-lg`}>
                  <stat.icon size={24} />
                </div>
                <div>
                  <div className="text-sm text-[#64748B] font-medium">{stat.label}</div>
                  <div className="text-2xl font-bold text-[#1E293B]">{stat.value}</div>
                </div>
              </div>
            </motion.div>
          ))}
        </div>

        <div className="grid grid-cols-1 lg:grid-cols-3 gap-8">
          {/* Recent Borrows Table */}
          <section className="lg:col-span-2">
            <div className="card bg-white overflow-hidden p-0">
              <div className="p-6 border-b border-[#E2E8F0] flex items-center justify-between">
                <h3 className="font-bold text-[#1E293B]">Transaksi Terbaru</h3>
                <Button variant="ghost" size="sm">Ekspor Data</Button>
              </div>
              <div className="overflow-x-auto">
                <table className="w-full text-left">
                  <thead className="bg-[#F8FAFC] text-[#64748B] text-xs font-bold uppercase tracking-wider">
                    <tr>
                      <th className="px-6 py-4">Anggota</th>
                      <th className="px-6 py-4">Buku</th>
                      <th className="px-6 py-4">Tgl Pinjam</th>
                      <th className="px-6 py-4">Status</th>
                      <th className="px-6 py-4 text-center">Aksi</th>
                    </tr>
                  </thead>
                  <tbody className="divide-y divide-[#E2E8F0]">
                    {borrows.slice(0, 5).map((b) => (
                      <tr key={b.id} className="hover:bg-[#F8FAFC] transition-colors">
                        <td className="px-6 py-4 font-medium text-[#1E293B]">{b.user.name}</td>
                        <td className="px-6 py-4 text-[#64748B]">{b.book.title}</td>
                        <td className="px-6 py-4 text-[#64748B]">{new Date(b.borrow_date).toLocaleDateString()}</td>
                        <td className="px-6 py-4">
                          <span className={`px-2 py-1 rounded-md text-[10px] font-bold uppercase ${
                            b.status === 'borrowed' ? 'bg-blue-100 text-blue-700' : 'bg-green-100 text-green-700'
                          }`}>
                            {b.status}
                          </span>
                        </td>
                        <td className="px-6 py-4 text-center">
                          <button className="text-[#64748B] hover:text-[#4338CA] mx-2"><Edit size={16} /></button>
                          <button className="text-[#64748B] hover:text-red-500 mx-2"><Trash2 size={16} /></button>
                        </td>
                      </tr>
                    ))}
                  </tbody>
                </table>
              </div>
            </div>
          </section>

          {/* User List Sideboard */}
          <section>
            <div className="card bg-white p-0 overflow-hidden">
               <div className="p-6 border-b border-[#E2E8F0]">
                <h3 className="font-bold text-[#1E293B]">Anggota Baru</h3>
              </div>
              <div className="divide-y divide-[#E2E8F0]">
                {users.slice(0, 6).map((u) => (
                  <div key={u.id} className="p-4 flex items-center justify-between hover:bg-[#F8FAFC] transition-colors group">
                    <div className="flex items-center gap-3">
                      <div className="w-10 h-10 bg-[#4338CA]/10 rounded-full flex items-center justify-center font-bold text-[#4338CA]">
                        {u.name.charAt(0)}
                      </div>
                      <div>
                        <div className="text-sm font-bold text-[#1E293B]">{u.name}</div>
                        <div className="text-xs text-[#64748B]">{u.email}</div>
                      </div>
                    </div>
                    <Button variant="ghost" size="sm" className="opacity-0 group-hover:opacity-100">Detail</Button>
                  </div>
                ))}
              </div>
            </div>
          </section>
        </div>
      </div>
    </main>
  );
}
