import { Navbar } from '@/components/layout/Navbar';
import { Button } from '@/components/ui/Button';
import { Search, BookOpen, Clock, ShieldCheck } from 'lucide-react';

export default function Home() {
  return (
    <main className="min-h-screen bg-[#F8FAFC]">
      <Navbar />
      
      {/* Hero Section */}
      <section className="pt-32 pb-20 px-4">
        <div className="max-w-7xl mx-auto text-center">
          <div className="inline-block px-4 py-1.5 mb-6 bg-[#4338CA]/10 text-[#4338CA] rounded-full text-sm font-semibold tracking-wide uppercase">
            Sistem Perpustakaan Modern
          </div>
          <h1 className="text-5xl md:text-6xl font-extrabold text-[#1E293B] mb-6 leading-tight">
            Pinjam Buku Kini Lebih <br />
            <span className="text-[#4338CA]">Mudah & Cepat</span>
          </h1>
          <p className="text-xl text-[#64748B] mb-10 max-w-2xl mx-auto">
            Akses ribuan katalog buku dari Open Library, Perpusnas, dan berbagai sumber lainnya dalam satu platform terintegrasi.
          </p>

          <div className="max-w-2xl mx-auto relative group">
            <div className="absolute inset-y-0 left-0 pl-4 flex items-center pointer-events-none">
              <Search className="h-5 w-5 text-[#94A3B8] group-focus-within:text-[#4338CA] transition-colors" />
            </div>
            <input
              type="text"
              placeholder="Cari judul buku, penulis, atau ISBN..."
              className="w-full pl-12 pr-4 py-4 bg-white border border-[#E2E8F0] rounded-2xl shadow-xl focus:outline-none focus:ring-2 focus:ring-[#4338CA]/20 focus:border-[#4338CA] transition-all text-lg"
            />
            <div className="absolute inset-y-2 right-2">
              <Button size="lg" className="h-full px-8 rounded-xl">
                Cari
              </Button>
            </div>
          </div>
        </div>
      </section>

      {/* Features Section */}
      <section className="py-20 bg-white">
        <div className="max-w-7xl mx-auto px-4 grid md:grid-cols-3 gap-12">
          <div className="text-center">
            <div className="w-16 h-16 bg-[#4338CA]/5 rounded-2xl flex items-center justify-center mx-auto mb-6">
              <BookOpen className="text-[#4338CA] w-8 h-8" />
            </div>
            <h3 className="text-xl font-bold mb-3 text-[#1E293B]">Katalog Lengkap</h3>
            <p className="text-[#64748B]">Akses ke ribuan judul buku dari berbagai kategori dan penerbit ternama.</p>
          </div>
          <div className="text-center">
            <div className="w-16 h-16 bg-[#14B8A6]/5 rounded-2xl flex items-center justify-center mx-auto mb-6">
              <Clock className="text-[#14B8A6] w-8 h-8" />
            </div>
            <h3 className="text-xl font-bold mb-3 text-[#1E293B]">Proses Instan</h3>
            <p className="text-[#64748B]">Lakukan peminjaman dan pengembalian buku secara mandiri hanya dalam hitungan detik.</p>
          </div>
          <div className="text-center">
            <div className="w-16 h-16 bg-[#4338CA]/5 rounded-2xl flex items-center justify-center mx-auto mb-6">
              <ShieldCheck className="text-[#4338CA] w-8 h-8" />
            </div>
            <h3 className="text-xl font-bold mb-3 text-[#1E293B]">Keamanan Terjamin</h3>
            <p className="text-[#64748B]">Sistem otentikasi JWT dan tracking peminjaman yang akurat untuk kenyamanan Anda.</p>
          </div>
        </div>
      </section>

      {/* Stats Section */}
      <section className="py-20 bg-[#F8FAFC]">
        <div className="max-w-7xl mx-auto px-4">
          <div className="bg-[#1E293B] rounded-[2rem] p-12 text-white relative overflow-hidden">
            <div className="absolute top-0 right-0 w-64 h-64 bg-[#4338CA]/20 rounded-full blur-[100px] -mr-32 -mt-32" />
            <div className="relative z-10 grid md:grid-cols-4 gap-8 text-center">
              <div>
                <div className="text-4xl font-bold mb-1">10k+</div>
                <div className="text-slate-400">Total Buku</div>
              </div>
              <div>
                <div className="text-4xl font-bold mb-1">2.5k+</div>
                <div className="text-slate-400">Anggota Aktif</div>
              </div>
              <div>
                <div className="text-4xl font-bold mb-1">50k+</div>
                <div className="text-slate-400">Transaksi</div>
              </div>
              <div>
                <div className="text-4xl font-bold mb-1">99%</div>
                <div className="text-slate-400">Kepuasan</div>
              </div>
            </div>
          </div>
        </div>
      </section>
      
      {/* Footer */}
      <footer className="py-12 border-t border-[#E2E8F0] text-center text-[#64748B] text-sm">
        &copy; 2026 Sistem Pinjam Perpus. Dibuat sesuai blueprint produksi.
      </footer>
    </main>
  );
}
