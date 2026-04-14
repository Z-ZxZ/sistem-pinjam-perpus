'use client';

import React, { useEffect, useState } from 'react';
import Image from 'next/image';
import { Navbar } from '@/components/layout/Navbar';
import { Button } from '@/components/ui/Button';
import { api } from '@/lib/api';
import { Search, Filter, Book as BookIcon } from 'lucide-react';
import toast from 'react-hot-toast';

interface Book {
  id: number;
  title: string;
  author: string;
  category: string;
  cover_url: string;
  stock: number;
}

import { useAuth } from '@/context/AuthContext';

export default function BookCatalog() {
  const { isLoggedIn } = useAuth();
  const [books, setBooks] = useState<Book[]>([]);
  const [search, setSearch] = useState('');
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    fetchBooks();
  }, []);

  const fetchBooks = async (val = '') => {
    setIsLoading(true);
    try {
      const res = await api.get(`/books?search=${val}&limit=20`);
      setBooks(res.books || []);
    } catch (err) {
      console.error(err);
    } finally {
      setIsLoading(false);
    }
  };

  const handleBorrow = async (bookId: number) => {
    if (!isLoggedIn) {
      toast.error('Silakan login terlebih dahulu untuk meminjam buku.');
      return;
    }

    try {
      await api.post('/borrow', { book_id: bookId });
      toast.success('Berhasil meminjam buku! Silakan cek dashboard Anda.');
      fetchBooks(search); // Biar stok nya ke refresh woi
    } catch (err: unknown) {
      if (err instanceof Error) {
        toast.error(err.message);
      } else {
        toast.error('Gagal meminjam buku. Silakan coba lagi.');
      }
    }
  };

  return (
    <main className="min-h-screen bg-[#F8FAFC]">
      <Navbar />
      
      <div className="max-w-7xl mx-auto px-4 pt-32 pb-20">
        <div className="flex flex-col md:flex-row md:items-center justify-between gap-6 mb-12">
          <div>
            <h1 className="text-3xl font-bold text-[#1E293B] mb-2">Katalog Buku</h1>
            <p className="text-[#64748B]">Temukan ribuan judul buku dari koleksi global kami.</p>
          </div>

          <div className="flex items-center gap-3 w-full md:w-auto">
            <div className="relative flex-1 md:w-80">
              <Search className="absolute left-3 top-1/2 -translate-y-1/2 text-[#94A3B8]" size={18} />
              <input
                type="text"
                placeholder="Cari buku..."
                className="w-full pl-10 pr-4 py-2 bg-white border border-[#E2E8F0] rounded-xl focus:ring-2 focus:ring-[#4338CA]/20 focus:outline-none transition-all"
                value={search}
                onChange={(e) => setSearch(e.target.value)}
                onKeyDown={(e) => e.key === 'Enter' && fetchBooks(search)}
              />
            </div>
            <Button variant="outline" className="flex items-center gap-2">
              <Filter size={18} />
              Filter
            </Button>
          </div>
        </div>

        {isLoading ? (
          <div className="grid grid-cols-1 md:grid-cols-4 lg:grid-cols-5 gap-6">
            {[...Array(10)].map((_, i) => (
              <div key={i} className="card h-80 animate-pulse bg-slate-100" />
            ))}
          </div>
        ) : (
          <div className="grid grid-cols-1 md:grid-cols-4 lg:grid-cols-5 gap-6">
            {books.map((book) => (
              <div key={book.id} className="card group hover:-translate-y-1 transition-all duration-300 p-0 overflow-hidden bg-white border-[#E2E8F0]">
                <div className="aspect-[3/4] bg-[#F1F5F9] relative overflow-hidden">
                  <div className="absolute inset-0 flex items-center justify-center">
                    <BookIcon size={48} className="text-[#CBD5E1] group-hover:scale-110 transition-transform duration-500" />
                  </div>
                  {book.cover_url && (
                    <Image 
                      src={book.cover_url} 
                      alt={book.title} 
                      fill 
                      className="object-cover"
                      unoptimized // Biarin unoptimized aja deh biar ga ribet config external URL wkwk
                    />
                  )}
                  <div className="absolute top-2 right-2">
                    <span className="bg-white/90 backdrop-blur-sm px-2 py-1 rounded-md text-[10px] font-bold text-[#4338CA] uppercase">
                      {book.category}
                    </span>
                  </div>
                </div>
                <div className="p-4">
                  <h3 className="font-bold text-[#1E293B] line-clamp-1 mb-1 group-hover:text-[#4338CA] transition-colors">{book.title}</h3>
                  <p className="text-xs text-[#64748B] mb-4">{book.author}</p>
                  <div className="flex items-center justify-between gap-2">
                     <span className={`text-[10px] font-bold ${book.stock > 0 ? 'text-green-600' : 'text-red-500'}`}>
                      {book.stock > 0 ? `${book.stock} Tersedia` : 'Habis'}
                    </span>
                    <Button 
                      size="sm" 
                      className="px-3 h-8 text-[10px]" 
                      disabled={book.stock === 0}
                      onClick={() => handleBorrow(book.id)}
                    >
                      Pinjam
                    </Button>
                  </div>
                </div>
              </div>
            ))}
          </div>
        )}
      </div>
    </main>
  );
}
