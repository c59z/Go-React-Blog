import { TextField, IconButton } from "@mui/material";
import SearchIcon from "@mui/icons-material/Search";
import CloseIcon from "@mui/icons-material/Close";
import "./index.scss";

interface SearchInputProps
  extends Omit<React.HTMLAttributes<HTMLDivElement>, "onChange"> {
  value: string;
  placeholder?: string;
  onChange: (val: string) => void;
  onSearch?: () => void;
  onClear?: () => void;
}

const SearchInput: React.FC<SearchInputProps> = ({
  value,
  placeholder = "Search...",
  onChange,
  onSearch,
  onClear,
  ...rest
}) => {
  return (
    <div className="search-input" {...rest}>
      <TextField
        size="small"
        value={value}
        placeholder={placeholder}
        onChange={(e) => onChange(e.target.value)}
        className="search-input__field"
      />

      {value && (
        <IconButton className="search-input__clear" onClick={onClear}>
          <CloseIcon fontSize="small" />
        </IconButton>
      )}

      <IconButton className="search-input__button" onClick={onSearch}>
        <SearchIcon />
      </IconButton>
    </div>
  );
};

export default SearchInput;
