import { Chip } from "@mui/material";
import "./index.scss";

export interface FilterOption {
  label: string;
  value: string;
}

interface FilterSectionProps {
  title: string;
  options: FilterOption[];
  value: string;
  onChange: (value: string) => void;
}

const FilterSection: React.FC<FilterSectionProps> = ({
  title,
  options,
  value,
  onChange,
}) => {
  return (
    <div className="filter-section">
      <div className="filter-section__title">{title}</div>

      <div className="filter-section__options">
        {options.map((item) => (
          <Chip
            key={item.value}
            label={item.label}
            clickable
            className={
              item.value === value
                ? "filter-chip filter-chip--active"
                : "filter-chip"
            }
            onClick={() => onChange(item.value)}
          />
        ))}
      </div>
    </div>
  );
};

export default FilterSection;
