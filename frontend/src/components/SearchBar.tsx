import { Search } from 'lucide-react';
import { Input } from '@/components/ui/input';

interface SearchBarProps {
  value: string;
  onChange: (value: string) => void;
}

const SearchBar = ({ value, onChange }: SearchBarProps) => {
  return (
    <div className="relative w-full max-w-3xl mx-auto group">
      <Search className="absolute left-5 top-1/2 -translate-y-1/2 h-6 w-6 text-primary group-focus-within:text-accent transition-colors duration-300" />
      <Input
        type="text"
        placeholder="Discover your next favorite meal..."
        value={value}
        onChange={(e) => onChange(e.target.value)}
        className="pl-14 h-16 text-xl border-white/10 focus:border-primary focus:ring-2 focus:ring-primary/20 bg-card/40 backdrop-blur-lg rounded-2xl shadow-neon transition-all"
      />
    </div>
  );
};

export default SearchBar;
