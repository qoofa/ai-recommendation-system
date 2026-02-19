import { useState, useMemo, useEffect } from 'react';
import FoodCard from '@/components/FoodCard';
import SearchBar from '@/components/SearchBar';
import heroImage from '/assets/hero-restaurant.jpg';

const Home = () => {
  const [searchQuery, setSearchQuery] = useState('');
  const [items, setItems] = useState([]);           // All food items
  const [searchResults, setSearchResults] = useState([]); // Semantic results
  const [loading, setLoading] = useState(true);

  // Load entire menu on first load
  useEffect(() => {
    async function fetchItems() {
      try {
        const res = await fetch("http://localhost:5000/api/v1/food");
        const data = await res.json();
        if (data.success) setItems(data.data);
      } catch (err) {
        console.error("Error fetching menu:", err);
      } finally {
        setLoading(false);
      }
    }
    fetchItems();
  }, []);

  useEffect(() => {
    const delayDebounce = setTimeout(async () => {
      if (!searchQuery.trim()) {
        setSearchResults([]);
        return;
      }

      try {
        const res = await fetch(
          `http://localhost:5000/api/v1/food/search?query=${encodeURIComponent(searchQuery)}`
        );
        const data = await res.json();
        if (data.success) {
          setSearchResults(data.data);
        }
      } catch (err) {
        console.error("Search error:", err);
      }
    }, 400);

    return () => clearTimeout(delayDebounce);
  }, [searchQuery]);

  const recommendations = useMemo(() => {
    return [...items]
      .sort((a, b) => (b.salesCount || 0) - (a.salesCount || 0))
      .slice(0, 3);
  }, [items]);

  if (loading) {
    return <div className="text-center py-12 text-xl">Loading menu...</div>;
  }

  const displayItems = searchQuery ? searchResults : items;

  return (
    <div className="min-h-screen mesh-bg">
      {/* Hero Section */}
      <section className="relative h-[85vh] flex items-center justify-center overflow-hidden">
        <div className="absolute inset-0 z-0">
          <img
            src={heroImage}
            alt="Restaurant hero"
            className="w-full h-full object-cover opacity-40 scale-110 blur-[2px]"
          />
          <div className="absolute inset-0 bg-gradient-to-b from-background/40 via-background/80 to-background" />
        </div>

        <div className="relative z-10 container text-center space-y-8 animate-float">
          <h1 className="text-7xl md:text-9xl font-black text-neon tracking-tighter filter drop-shadow-[0_0_30px_hsl(var(--primary)/0.5)]">
            MFC
          </h1>
          <p className="text-2xl md:text-3xl text-foreground font-light max-w-3xl mx-auto tracking-wide">
            The Future of <span className="text-accent font-bold">Culinary Art</span>
          </p>
          <div className="flex justify-center gap-4 pt-6">
            <div className="h-1 w-20 bg-primary rounded-full animate-pulse-neon" />
            <div className="h-1 w-20 bg-accent rounded-full animate-pulse-neon delay-75" />
          </div>
        </div>
      </section>

      {/* Search */}
      <section className="container pt-12">
        <SearchBar value={searchQuery} onChange={setSearchQuery} />
      </section>

      {/* Recommendations */}
      {!searchQuery && (
        <section className="container pb-12 pt-6">
          <div className="space-y-10">
            <div className="text-center space-y-3">
              <h2 className="text-4xl md:text-5xl font-black text-neon">Chef's Matrix</h2>
              <p className="text-accent font-bold text-lg tracking-widest uppercase">High Performance Dishes</p>
            </div>

            <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
              {recommendations.map(item => (
                <FoodCard key={item.id} item={item} />
              ))}
            </div>
          </div>
        </section>
      )}

      {/* Menu / Search Results */}
      <section className="container pt-12 pb-24">
        <div className="space-y-12">
          <div className="text-center space-y-4">
            <h2 className="text-4xl md:text-6xl font-black tracking-tighter">
              {searchQuery ? 'Grid Results' : 'System Menu'}
            </h2>

            {!searchQuery && <p className="text-muted-foreground text-xl font-light">Precision crafted. Served with intent.</p>}
          </div>

          {displayItems.length > 0 ? (
            <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
              {displayItems.map(item => (
                <FoodCard key={item.id} item={item} />
              ))}
            </div>
          ) : (
            <div className="text-center py-12">
              <p className="text-xl text-muted-foreground">
                No dishes found matching "{searchQuery}"
              </p>
            </div>
          )}
        </div>
      </section>
    </div>
  );
};

export default Home;
