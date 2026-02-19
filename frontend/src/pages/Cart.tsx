import { useCart } from '@/context/CartContext';
import { Button } from '@/components/ui/button';
import { Card, CardContent, CardFooter, CardHeader, CardTitle } from '@/components/ui/card';
import { Minus, Plus, Trash2, ShoppingBag } from 'lucide-react';
import { useNavigate } from 'react-router-dom';

const Cart = () => {
  const { cart, removeFromCart, updateQuantity, getTotalPrice } = useCart();
  const navigate = useNavigate();

  if (cart.length === 0) {
    return (
      <div className="container py-24 text-center">
        <div className="max-w-md mx-auto space-y-6">
          <ShoppingBag className="h-24 w-24 mx-auto text-muted-foreground" />
          <h2 className="text-3xl font-bold">Your cart is empty</h2>
          <p className="text-muted-foreground">
            Start adding some delicious dishes to your cart
          </p>
          <Button onClick={() => navigate('/')} size="lg">
            Browse Menu
          </Button>
        </div>
      </div>
    );
  }

  return (
    <div className="min-h-screen mesh-bg py-20 px-4">
      <div className="container max-w-6xl mx-auto">
        <h1 className="text-6xl font-black mb-12 text-neon tracking-tighter">Your Command</h1>
        
        <div className="grid lg:grid-cols-3 gap-12">
          <div className="lg:col-span-2 space-y-4">
            {cart.map(item => (
              <Card key={item.id} className="glass-card border-white/5 hover-lift overflow-hidden">
                <CardContent className="p-0">
                  <div className="flex flex-col md:flex-row gap-0">
                    <div className="w-full md:w-32 h-32 relative">
                      <img
                        src={"/assets/"+item.image}
                        alt={item.name}
                        className="w-full h-full object-cover"
                      />
                      <div className="absolute inset-0 bg-primary/10" />
                    </div>
                    
                    <div className="flex-1 p-6 space-y-4">
                      <div className="flex justify-between items-start">
                        <div>
                          <h3 className="font-black text-2xl tracking-tight">{item.name}</h3>
                          <p className="text-sm text-muted-foreground font-light line-clamp-1">
                            {item.description}
                          </p>
                        </div>
                        <Button
                          variant="ghost"
                          size="icon"
                          onClick={() => removeFromCart(item.id)}
                          className="text-destructive hover:bg-destructive/10 hover:text-destructive w-10 h-10"
                        >
                          <Trash2 className="h-5 w-5" />
                        </Button>
                      </div>
                      
                      <div className="flex justify-between items-center">
                        <div className="flex items-center gap-4 bg-background/40 p-1 rounded-xl border border-white/5">
                          <Button
                            variant="ghost"
                            size="icon"
                            onClick={() => updateQuantity(item.id, item.quantity - 1)}
                            className="h-10 w-10 hover:bg-primary/20"
                          >
                            <Minus className="h-4 w-4" />
                          </Button>
                          <span className="font-black text-lg w-10 text-center">
                            {item.quantity}
                          </span>
                          <Button
                            variant="ghost"
                            size="icon"
                            onClick={() => updateQuantity(item.id, item.quantity + 1)}
                            className="h-10 w-10 hover:bg-primary/20"
                          >
                            <Plus className="h-4 w-4" />
                          </Button>
                        </div>
                        
                        <p className="text-2xl font-black text-accent">
                          ${(item.price * item.quantity).toFixed(2)}
                        </p>
                      </div>
                    </div>
                  </div>
                </CardContent>
              </Card>
            ))}
          </div>
          
          <div className="lg:col-span-1">
            <Card className="glass-card border-white/5 sticky top-28 overflow-hidden">
              <div className="h-2 bg-gradient-to-r from-primary via-accent to-purple-500" />
              <CardHeader>
                <CardTitle className="text-2xl font-black">Extraction Summary</CardTitle>
              </CardHeader>
              <CardContent className="space-y-6">
                <div className="flex justify-between text-lg font-light">
                  <span>Core Value</span>
                  <span>${getTotalPrice().toFixed(2)}</span>
                </div>
                <div className="flex justify-between text-lg font-light">
                  <span>System Tax (10%)</span>
                  <span>${(getTotalPrice() * 0.1).toFixed(2)}</span>
                </div>
                <div className="border-t border-white/10 pt-6">
                  <div className="flex justify-between text-3xl font-black">
                    <span>Total</span>
                    <span className="text-neon">
                      ${(getTotalPrice() * 1.1).toFixed(2)}
                    </span>
                  </div>
                </div>
              </CardContent>
              <CardFooter className="flex flex-col gap-4 pb-8">
                <Button 
                  className="w-full bg-primary hover:bg-primary/80 h-14 text-lg font-black neon-glow transition-all" 
                  size="lg"
                  onClick={() => navigate('/checkout')}
                >
                  Initialize Protocol
                </Button>
                <Button 
                  variant="outline" 
                  className="w-full border-white/10 h-14 text-lg font-bold hover:bg-white/5"
                  onClick={() => navigate('/')}
                >
                  Abort & Return
                </Button>
              </CardFooter>
            </Card>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Cart;
