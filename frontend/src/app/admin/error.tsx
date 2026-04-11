'use client';

import { useEffect } from 'react';
import { Button } from '@/components/ui/Button';

export default function Error({
  error,
  reset,
}: {
  error: Error & { digest?: string };
  reset: () => void;
}) {
  useEffect(() => {
    console.error('Page Error:', error);
  }, [error]);

  return (
    <div className="min-h-screen bg-[#F8FAFC] flex flex-col items-center justify-center p-4">
      <div className="bg-white p-8 rounded-2xl shadow-xl border border-[#E2E8F0] max-w-md w-full text-center">
        <div className="w-16 h-16 bg-red-100 text-red-600 rounded-full flex items-center justify-center mx-auto mb-4 text-2xl font-bold">
          !
        </div>
        <h2 className="text-2xl font-bold text-[#1E293B] mb-2">Terjadi Kesalahan</h2>
        <p className="text-[#64748B] mb-6">
          {error.message || 'Terjadi kesalahan saat memuat halaman ini.'}
        </p>
        <div className="flex flex-col gap-3">
          <Button onClick={() => reset()} className="w-full">
            Coba Lagi
          </Button>
          <Button variant="outline" onClick={() => window.location.href = '/'} className="w-full">
            Kembali ke Beranda
          </Button>
        </div>
      </div>
    </div>
  );
}
