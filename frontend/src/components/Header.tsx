import { ShoppingCart } from 'lucide-react';
import { Button } from '@/components/ui/button';
import { useCart } from '@/context/CartContext';
import { useNavigate, useLocation } from 'react-router-dom';

const Header = () => {
  const { getTotalItems } = useCart();
  const navigate = useNavigate();
  const location = useLocation();
  const totalItems = getTotalItems();

  return (
    <header className="sticky top-0 z-50 w-full border-b border-white/5 bg-background/60 backdrop-blur-xl supports-[backdrop-filter]:bg-background/40">
      <div className="container flex h-20 items-center justify-between">
        <button 
          onClick={() => navigate('/')}
          className="text-3xl font-black text-neon cursor-pointer hover:scale-105 transition-transform duration-300 tracking-tighter"
        >
          MFC
        </button>
        
        {location.pathname !== '/cart' && location.pathname !== '/checkout' && (
          <Button
            variant="outline"
            size="icon"
            className="relative neon-glow border-primary/40 bg-card/50 hover-lift h-12 w-12"
            onClick={() => navigate('/cart')}
          >
            <ShoppingCart className="h-6 w-6" />
            {totalItems > 0 && (
              <span className="absolute -top-3 -right-3 h-6 w-6 rounded-full bg-primary text-primary-foreground text-[10px] flex items-center justify-center font-bold animate-pulse-neon shadow-lg">
                {totalItems}
              </span>
            )}
          </Button>
        )}
      </div>
    </header>
  );
};

export default Header;
