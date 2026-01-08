import { useState } from "react";
import { TextField, Popper, Paper, MenuItem, Chip } from "@mui/material";
import "./index.scss";

type Mode = "idle" | "text" | "advanced";

type KeywordType = "tag" | "category";

interface KeywordConfig {
  prefix: string;
  color: "info" | "success";
  label: string;
}

const KEYWORDS: Record<KeywordType, KeywordConfig> = {
  tag: {
    prefix: "$tag:",
    color: "info",
    label: "标签",
  },
  category: {
    prefix: "$category:",
    color: "success",
    label: "分类",
  },
};

export const SearchBar = () => {
  const [anchorEl, setAnchorEl] = useState<HTMLElement | null>(null);
  const [focus, setFocus] = useState(false);

  const [input, setInput] = useState("");
  const [tags, setTags] = useState<string[]>([]);

  const trimmed = input.trim();

  const matchedKeywordEntry = Object.entries(KEYWORDS).find(([, cfg]) =>
    trimmed.startsWith(cfg.prefix)
  );

  const matchedType = matchedKeywordEntry
    ? (matchedKeywordEntry[0] as KeywordType)
    : null;

  const matchedConfig = matchedType ? KEYWORDS[matchedType] : null;

  const mode: Mode = matchedConfig
    ? "advanced"
    : trimmed === ""
    ? "idle"
    : "text";

  const keywordValue = matchedConfig
    ? trimmed.slice(matchedConfig.prefix.length).trim()
    : "";

  const commitValue = (value: string) => {
    if (!value) return;
    setTags((prev) => [...prev, value]);
    setInput("");
  };

  const commitText = (value: string, exclude = false) => {
    commitValue(exclude ? `-${value}` : value);
  };

  const commitKeyword = (type: KeywordType, value: string, exclude = false) => {
    if (!value) return;
    const prefix = KEYWORDS[type].prefix;
    commitValue(`${exclude ? "-" : ""}${prefix}${value}`);
  };

  const handleKeyDown = (e: React.KeyboardEvent<HTMLInputElement>) => {
    if (e.key !== "Enter") return;
    e.preventDefault();

    if (mode === "text") {
      commitText(trimmed);
    }

    if (mode === "advanced" && matchedType) {
      commitKeyword(matchedType, keywordValue);
    }
  };

  const renderChip = (value: string, index: number) => {
    let color: "default" | "primary" | "info" | "success" | "error" = "primary";

    const isExclude = value.startsWith("-");
    const raw = isExclude ? value.slice(1) : value;

    for (const cfg of Object.values(KEYWORDS)) {
      if (raw.startsWith(cfg.prefix)) {
        color = isExclude ? "error" : cfg.color;
        break;
      }
    }

    if (isExclude && color === "primary") {
      color = "error";
    }

    return (
      <Chip
        key={index}
        label={value}
        size="small"
        color={color}
        onDelete={() => setTags((prev) => prev.filter((_, i) => i !== index))}
      />
    );
  };

  return (
    <div className="search-container">
      <TextField
        className="search-bar"
        fullWidth
        placeholder="搜索"
        value={input}
        onFocus={(e) => {
          setFocus(true);
          setAnchorEl(e.currentTarget);
        }}
        onBlur={() => setTimeout(() => setFocus(false), 150)}
        onChange={(e) => setInput(e.target.value)}
        onKeyDown={handleKeyDown}
        slotProps={{
          input: {
            startAdornment: (
              <div className="tag-area">{tags.map(renderChip)}</div>
            ),
          },
        }}
      />

      <Popper
        open={focus}
        anchorEl={anchorEl}
        placement="bottom-start"
        className="search-popper"
      >
        <Paper className="search-dropdown">
          {/* idle */}
          {mode === "idle" && (
            <>
              {Object.entries(KEYWORDS).map(([type, cfg]) => (
                <MenuItem
                  key={type}
                  onMouseDown={(e) => e.preventDefault()}
                  onClick={() => setInput(cfg.prefix)}
                >
                  🔍 搜索{cfg.label}（{cfg.prefix}）
                </MenuItem>
              ))}
            </>
          )}

          {/* text */}
          {mode === "text" && (
            <>
              <MenuItem onClick={() => commitText(trimmed)}>
                🔍 “{trimmed}” 纯文本搜索
              </MenuItem>
              <MenuItem onClick={() => commitText(trimmed, true)}>
                🚫 -“{trimmed}” 排除搜索
              </MenuItem>
            </>
          )}

          {/* advanced */}
          {mode === "advanced" && matchedConfig && matchedType && (
            <>
              <MenuItem
                onClick={() => commitKeyword(matchedType, keywordValue)}
              >
                🔍 搜索{matchedConfig.label} ({matchedConfig.prefix})：
                {keywordValue || "（未填写）"}
              </MenuItem>

              <MenuItem
                onClick={() => commitKeyword(matchedType, keywordValue, true)}
              >
                🚫 排除{matchedConfig.label} ({matchedConfig.prefix})：
                {keywordValue || "（未填写）"}
              </MenuItem>
            </>
          )}
        </Paper>
      </Popper>
    </div>
  );
};
