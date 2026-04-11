'use client';

import React, { useEffect, useState } from 'react';
import { Navbar } from '@/components/layout/Navbar';
import { Button } from '@/components/ui/Button';
import { api } from '@/lib/api';
import { Book, Clock, AlertCircle, BookOpen, ChevronRight } from 'lucide-react';
import { motion } from 'framer-motion';

interface User {
  id: number;
  name: string;
  email: string;
}

interface Book {
  id: number;
  title: string;
}

interface Borrow {
  id: number;
  book: Book;
  due_date: string;
  status: string;
}

interface Fine {
  id: number;
  amount: number;
  created_at: string;
}

export default function Dashboard() {
  const [history, setHistory] = useState<Borrow[]>([]);
  const [fines, setFines] = useState<{ total: number; fines: Fine[] }>({ total: 0, fines: [] });
  const [user, setUser] = useState<User | null>(null);
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    const userData = localStorage.getItem('user');
    if (userData) setUser(JSON.parse(userData));

    const fetchData = async () => {
      try {
        const [historyRes, finesRes] = await Promise.all([
          api.get('/history'),
          api.get('/fines'),
        ]);
        setHistory(historyRes.data);
        setFines(finesRes.data);
      } catch (err) {
        console.error(err);
      } finally {
        setIsLoading(false);
      }
    };

    fetchData();
  }, []);

  if (isLoading) return <div className="min-h-screen bg-[#F8FAFC]" />;

  return (
    <main className="min-h-screen bg-[#F8FAFC] pb-20">
      <Navbar />
      
      <div className="max-w-7xl mx-auto px-4 pt-32">
        <div className="mb-10">
          <h1 className="text-3xl font-bold text-[#1E293B] mb-2">Halo, {user?.name}</h1>
          <p className="text-[#64748B]">Berikut adalah ringkasan aktivitas perpustakaan Anda.</p>
        </div>

        <div className="grid grid-cols-1 md:grid-cols-3 gap-6 mb-12">
          {/* Stats Cards */}
          <motion.div initial={{ opacity: 0, y: 20 }} animate={{ opacity: 1, y: 0 }} className="card bg-white">
            <div className="flex items-center gap-4">
              <div className="w-12 h-12 bg-[#4338CA]/10 rounded-xl flex items-center justify-center">
                <BookOpen className="text-[#4338CA]" />
              </div>
              <div>
                <div className="text-sm text-[#64748B] font-medium">Buku Dipinjam</div>
                <div className="text-2xl font-bold text-[#1E293B]">{history.filter(h => h.status === 'borrowed').length}</div>
              </div>
            </div>
          </motion.div>

          <motion.div initial={{ opacity: 0, y: 20 }} animate={{ opacity: 1, y: 0 }} transition={{ delay: 0.1 }} className="card bg-white">
            <div className="flex items-center gap-4">
              <div className="w-12 h-12 bg-[#14B8A6]/10 rounded-xl flex items-center justify-center">
                <Clock className="text-[#14B8A6]" />
              </div>
              <div>
                <div className="text-sm text-[#64748B] font-medium">Jatuh Tempo</div>
                <div className="text-2xl font-bold text-[#1E293B]">{history.filter(h => h.status === 'overdue').length}</div>
              </div>
            </div>
          </motion.div>

          <motion.div initial={{ opacity: 0, y: 20 }} animate={{ opacity: 1, y: 0 }} transition={{ delay: 0.2 }} className="card bg-white">
            <div className="flex items-center gap-4">
              <div className="w-12 h-12 bg-red-50 rounded-xl flex items-center justify-center">
                <AlertCircle className="text-red-600" />
              </div>
              <div>
                <div className="text-sm text-[#64748B] font-medium">Total Denda</div>
                <div className="text-2xl font-bold text-red-600">Rp {fines.total.toLocaleString()}</div>
              </div>
            </div>
          </motion.div>
        </div>

        <div className="grid grid-cols-1 lg:grid-cols-2 gap-8">
          {/* Borrowing History */}
          <section>
            <div className="flex items-center justify-between mb-6">
              <h2 className="text-xl font-bold text-[#1E293B]">Riwayat Peminjaman</h2>
              <Button variant="ghost" size="sm">Lihat Semua</Button>
            </div>
            
            <div className="space-y-4">
              {history.length === 0 ? (
                <div className="card text-center py-12 text-[#64748B]">Belum ada riwayat peminjaman.</div>
              ) : (
                history.map((item) => (
                  <div key={item.id} className="card bg-white flex items-center justify-between group">
                    <div className="flex items-center gap-4">
                      <div className="w-12 h-16 bg-[#F8FAFC] rounded-lg border border-[#E2E8F0] flex items-center justify-center">
                        <Book className="text-[#94A3B8] w-6 h-6" />
                      </div>
                      <div>
                        <h4 className="font-bold text-[#1E293B] group-hover:text-[#4338CA] transition-colors">{item.book.title}</h4>
                        <p className="text-sm text-[#64748B]">Batas Kembali: {new Date(item.due_date).toLocaleDateString('id-ID')}</p>
                      </div>
                    </div>
                    <div className="flex items-center gap-4">
                      <span className={`px-3 py-1 rounded-full text-xs font-bold uppercase ${
                        item.status === 'borrowed' ? 'bg-blue-100 text-blue-700' :
                        item.status === 'returned' ? 'bg-green-100 text-green-700' :
                        'bg-red-100 text-red-700'
                      }`}>
                        {item.status}
                      </span>
                      <ChevronRight className="text-[#94A3B8]" />
                    </div>
                  </div>
                ))
              )}
            </div>
          </section>

          {/* Fines Activity */}
          <section>
            <div className="flex items-center justify-between mb-6">
              <h2 className="text-xl font-bold text-[#1E293B]">Denda Aktif</h2>
            </div>
            
            <div className="space-y-4">
              {fines.fines.length === 0 ? (
                <div className="card bg-[#14B8A6]/5 border-[#14B8A6]/20 text-center py-12 text-[#0D9488]">
                  Hebat! Anda tidak memiliki denda tertunggak.
                </div>
              ) : (
                fines.fines.map((fine: Fine) => (
                  <div key={fine.id} className="card bg-white border-l-4 border-red-500 py-4">
                    <div className="flex justify-between items-center">
                      <div>
                        <div className="text-sm font-semibold text-[#1E293B]">Terlambat Pengembalian</div>
                        <div className="text-xs text-[#64748B]">{new Date(fine.created_at).toLocaleDateString('id-ID')}</div>
                      </div>
                      <div className="text-lg font-bold text-red-600">Rp {fine.amount.toLocaleString()}</div>
                    </div>
                  </div>
                ))
              )}
            </div>
          </section>
        </div>
      </div>
    </main>
  );
}
