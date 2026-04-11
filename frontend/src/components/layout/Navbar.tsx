'use client';

import React from 'react';
import Link from 'next/link';
import { Button } from '@/components/ui/Button';
import { Book, LogOut, LayoutDashboard } from 'lucide-react';

export const Navbar = () => {
  // We'll add real auth logic later
  const isLoggedIn = false;
  const isAdmin = false;

  return (
    <nav className="fixed top-0 left-0 right-0 h-16 bg-white/80 backdrop-blur-md border-b border-[#E2E8F0] z-50">
      <div className="max-w-7xl mx-auto px-4 h-full flex items-center justify-between">
        <Link href="/" className="flex items-center gap-2">
          <div className="w-8 h-8 bg-[#4338CA] rounded-lg flex items-center justify-center">
            <Book className="text-white w-5 h-5" />
          </div>
          <span className="font-bold text-xl tracking-tight text-[#1E293B]">
            PerpusDigital
          </span>
        </Link>

        <div className="flex items-center gap-6">
          <Link href="/books" className="text-[#64748B] hover:text-[#4338CA] transition-colors font-medium">
            Katalog
          </Link>
          
          {isLoggedIn ? (
            <>
              <Link href="/dashboard" className="text-[#64748B] hover:text-[#4338CA] transition-colors font-medium">
                Peminjaman
              </Link>
              {isAdmin && (
                <Link href="/admin" className="text-[#64748B] hover:text-[#4338CA] transition-colors font-medium flex items-center gap-1">
                  <LayoutDashboard size={18} />
                  Admin
                </Link>
              )}
              <div className="h-6 w-[1px] bg-[#E2E8F0]" />
              <button className="text-[#64748B] hover:text-red-600 transition-colors">
                <LogOut size={20} />
              </button>
            </>
          ) : (
            <div className="flex items-center gap-2">
              <Link href="/login">
                <Button variant="ghost">Masuk</Button>
              </Link>
              <Link href="/register">
                <Button>Daftar</Button>
              </Link>
            </div>
          )}
        </div>
      </div>
    </nav>
  );
};
