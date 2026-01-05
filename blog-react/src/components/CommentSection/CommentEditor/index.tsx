import { useState } from "react";
import "./index.scss";
import { CircularProgress } from "@mui/material";

interface Props {
  onSubmit: (content: string) => void;
  loading: boolean;
}

const CommentEditor = ({ onSubmit, loading }: Props) => {
  const [value, setValue] = useState("");
  const [focused, setFocused] = useState(false);

  const handleCancel = () => {
    setValue("");
    setFocused(false);
  };

  const handleSubmit = () => {
    if (!value.trim()) return;
    onSubmit(value.trim());
    setValue("");
    setFocused(false);
  };

  return (
    <div className={`comment-editor ${focused ? "active" : ""}`}>
      <textarea
        placeholder="发表评论..."
        value={value}
        onChange={(e) => setValue(e.target.value)}
        onFocus={() => setFocused(true)}
        rows={focused ? 3 : 1}
      />

      {focused && (
        <div className="editor-actions">
          <button className="btn-cancel" onClick={handleCancel}>
            取消
          </button>
          <button
            className="btn-submit"
            disabled={!value.trim() || loading}
            onClick={handleSubmit}
          >
            {loading ? <CircularProgress size={16} /> : "发表评论"}
          </button>
        </div>
      )}
    </div>
  );
};

export default CommentEditor;
