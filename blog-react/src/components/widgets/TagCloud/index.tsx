import { Box, Typography, CircularProgress } from "@mui/material";
import "./index.scss";
import { useNavigate } from "react-router-dom";
import { useArticleTags } from "./hook/useArticleTags";

const TagCloud = () => {
  const navigate = useNavigate();
  const { loading, tags } = useArticleTags();

  return (
    <Box className="tag-cloud-root">
      <Typography className="tag-cloud-title">Tags</Typography>

      {loading ? (
        <CircularProgress size={24} />
      ) : (
        <Box className="tag-cloud-list">
          {tags.map((item) => (
            <span
              key={item.tag}
              className="tag-cloud-item"
              data-count={item.number}
              onClick={() => navigate(`/search?tag=${item.tag}`)}
            >
              #{item.tag}
            </span>
          ))}
        </Box>
      )}
    </Box>
  );
};

export default TagCloud;
