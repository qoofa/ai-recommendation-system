import { useState } from 'react';
import { useCart } from '@/context/CartContext';
import { Button } from '@/components/ui/button';
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card';
import { Input } from '@/components/ui/input';
import { Label } from '@/components/ui/label';
import { useNavigate } from 'react-router-dom';
import { toast } from '@/hooks/use-toast';
import { CreditCard, MapPin, User } from 'lucide-react';
import axios from 'axios';

const Checkout = () => {
  const { cart, getTotalPrice, clearCart } = useCart();
  const navigate = useNavigate();
  const [isProcessing, setIsProcessing] = useState(false);

  const [formData, setFormData] = useState({
    fullName: '',
    email: '',
    phone: '',
    address: '',
    city: '',
    zipCode: '',
    cardNumber: '',
    cardExpiry: '',
    cardCvv: '',
  });

  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    setFormData(prev => ({ ...prev, [name]: value }));
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    // BASIC VALIDATION
    const requiredFields = Object.entries(formData);
    const emptyField = requiredFields.find(([_, value]) => !value.trim());

    if (emptyField) {
      toast({
        title: "Incomplete form",
        description: "Please fill in all fields",
        variant: "destructive",
      });
      return;
    }

    setIsProcessing(true);

    try {
      const itemIds = cart.map(item => item.id);
      console.log("🚀 ~ handleSubmit ~ cart:", cart)
      console.log("Training combo with:", itemIds);

      await axios.post("http://localhost:5000/api/v1/order/train", {
        items: itemIds,
      });

      console.log("Training API hit successfully!");

    } catch (err) {
      console.error("Training API error:", err);
      toast({
        title: "AI training failed",
        description: "Order placed, but combo learning couldn't update.",
        variant: "destructive",
      });
    }

    // ----------------------------------------------------
    // 🔥 SIMULATE PAYMENT SUCCESS
    // ----------------------------------------------------
    setTimeout(() => {
      setIsProcessing(false);

      clearCart();

      toast({
        title: "Order placed successfully!",
        description: `Your order of $${(getTotalPrice() * 1.1).toFixed(2)} has been confirmed`,
      });

      navigate('/');
    }, 2000);
  };

  if (cart.length === 0) {
    navigate('/cart');
    return null;
  }

  const totalAmount = (getTotalPrice() * 1.1).toFixed(2);

  return (
    <div className="min-h-screen mesh-bg py-20 px-4">
      <div className="container max-w-6xl mx-auto">
        <h1 className="text-6xl font-black mb-12 text-neon tracking-tighter">Secure Protocol</h1>

        <div className="grid lg:grid-cols-3 gap-12">
          
          {/* LEFT SIDE FORM */}
          <div className="lg:col-span-2">
            <form onSubmit={handleSubmit} className="space-y-6">

              {/* Personal Info */}
              <Card className="glass-card border-white/5 overflow-hidden">
                <CardHeader className="bg-primary/5">
                  <CardTitle className="flex items-center gap-3 text-2xl font-black">
                    <User className="h-6 w-6 text-primary" />
                    Operator Profile
                  </CardTitle>
                </CardHeader>

                <CardContent className="space-y-4">
                  <div className="grid md:grid-cols-2 gap-4">
                    <div className="space-y-2">
                      <Label htmlFor="fullName">Full Name</Label>
                      <Input
                        id="fullName"
                        name="fullName"
                        value={formData.fullName}
                        onChange={handleInputChange}
                        placeholder="John Doe"
                      />
                    </div>

                    <div className="space-y-2">
                      <Label htmlFor="email">Email</Label>
                      <Input
                        id="email"
                        name="email"
                        type="email"
                        value={formData.email}
                        onChange={handleInputChange}
                        placeholder="john@example.com"
                      />
                    </div>
                  </div>

                  <div className="space-y-2">
                    <Label htmlFor="phone">Phone Number</Label>
                    <Input
                      id="phone"
                      name="phone"
                      type="tel"
                      value={formData.phone}
                      onChange={handleInputChange}
                      placeholder="+1 (555) 123-4567"
                    />
                  </div>
                </CardContent>
              </Card>

              {/* Delivery Address */}
              <Card className="glass-card border-white/5 overflow-hidden">
                <CardHeader className="bg-accent/5">
                  <CardTitle className="flex items-center gap-3 text-2xl font-black">
                    <MapPin className="h-6 w-6 text-accent" />
                    Drop Point
                  </CardTitle>
                </CardHeader>

                <CardContent className="space-y-4">
                  <div className="space-y-2">
                    <Label htmlFor="address">Street Address</Label>
                    <Input
                      id="address"
                      name="address"
                      value={formData.address}
                      onChange={handleInputChange}
                      placeholder="123 Main St"
                    />
                  </div>

                  <div className="grid md:grid-cols-2 gap-4">
                    <div className="space-y-2">
                      <Label htmlFor="city">City</Label>
                      <Input
                        id="city"
                        name="city"
                        value={formData.city}
                        onChange={handleInputChange}
                        placeholder="New York"
                      />
                    </div>

                    <div className="space-y-2">
                      <Label htmlFor="zipCode">ZIP Code</Label>
                      <Input
                        id="zipCode"
                        name="zipCode"
                        value={formData.zipCode}
                        onChange={handleInputChange}
                        placeholder="10001"
                      />
                    </div>
                  </div>
                </CardContent>
              </Card>

              {/* Payment Info */}
              <Card className="glass-card border-white/5 overflow-hidden">
                <CardHeader className="bg-purple-500/5">
                  <CardTitle className="flex items-center gap-3 text-2xl font-black">
                    <CreditCard className="h-6 w-6 text-purple-400" />
                    Credit Transfer
                  </CardTitle>
                  <CardDescription className="text-purple-300/60">
                    Encrypted secure transmission - Simulation only
                  </CardDescription>
                </CardHeader>

                <CardContent className="space-y-4">
                  <div className="space-y-2">
                    <Label htmlFor="cardNumber">Card Number</Label>
                    <Input
                      id="cardNumber"
                      name="cardNumber"
                      value={formData.cardNumber}
                      onChange={handleInputChange}
                      placeholder="1234 5678 9012 3456"
                      maxLength={19}
                    />
                  </div>

                  <div className="grid md:grid-cols-2 gap-4">
                    <div className="space-y-2">
                      <Label htmlFor="cardExpiry">Expiry Date</Label>
                      <Input
                        id="cardExpiry"
                        name="cardExpiry"
                        value={formData.cardExpiry}
                        onChange={handleInputChange}
                        placeholder="MM/YY"
                        maxLength={5}
                      />
                    </div>

                    <div className="space-y-2">
                      <Label htmlFor="cardCvv">CVV</Label>
                      <Input
                        id="cardCvv"
                        name="cardCvv"
                        type="password"
                        value={formData.cardCvv}
                        onChange={handleInputChange}
                        placeholder="123"
                        maxLength={4}
                      />
                    </div>
                  </div>
                </CardContent>
              </Card>

              {/* Submit Button */}
              <Button 
                type="submit" 
                size="lg" 
                className="w-full bg-primary hover:bg-primary/80 h-16 text-xl font-black neon-glow transition-all active:scale-[0.98]"
                disabled={isProcessing}
              >
                {isProcessing ? 'Processing...' : `Authorize Transfer - $${totalAmount}`}
              </Button>

            </form>
          </div>

          {/* ORDER SUMMARY */}
          <div className="lg:col-span-1">
            <Card className="glass-card border-white/5 sticky top-28 overflow-hidden">
              <div className="h-2 bg-gradient-to-r from-primary via-accent to-purple-500" />
              <CardHeader>
                <CardTitle className="text-2xl font-black">Manifest Summary</CardTitle>
              </CardHeader>

              <CardContent className="space-y-4">
                <div className="space-y-3">
                  {cart.map(item => (
                    <div key={item.id} className="flex justify-between text-sm">
                      <span className="text-muted-foreground">
                        {item.name} × {item.quantity}
                      </span>
                      <span>${(item.price * item.quantity).toFixed(2)}</span>
                    </div>
                  ))}
                </div>

                <div className="border-t border-border pt-4 space-y-2">
                  <div className="flex justify-between text-sm text-muted-foreground">
                    <span>Subtotal</span>
                    <span>${getTotalPrice().toFixed(2)}</span>
                  </div>

                  <div className="flex justify-between text-sm text-muted-foreground">
                    <span>Tax (10%)</span>
                    <span>${(getTotalPrice() * 0.1).toFixed(2)}</span>
                  </div>

                  <div className="flex justify-between text-2xl font-black pt-4">
                    <span>Net Total</span>
                    <span className="text-neon">${totalAmount}</span>
                  </div>
                </div>

              </CardContent>
            </Card>
          </div>

        </div>
      </div>
    </div>
  );
};

export default Checkout;
