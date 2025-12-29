import { Stack, Typography, Chip, Box } from "@mui/material";
import VisibilityOutlinedIcon from "@mui/icons-material/VisibilityOutlined";
import FavoriteBorderOutlinedIcon from "@mui/icons-material/FavoriteBorderOutlined";
import ChatBubbleOutlineOutlinedIcon from "@mui/icons-material/ChatBubbleOutlineOutlined";
import type { Article } from "@/api/article";
import "./index.scss";

interface Props {
  article: Article;
}

const ArticleHeader: React.FC<Props> = ({ article }) => {
  return (
    <div className="article-header">
      <Stack spacing={2}>
        <Typography variant="h4" fontWeight={600}>
          {article?.title}
        </Typography>

        <Typography variant="body1" color="text.secondary">
          {article?.abstract}
        </Typography>

        <Stack direction="row" spacing={1} flexWrap="wrap">
          <Chip label={article?.category} size="small" />
          <Chip label={article?.keyword} size="small" variant="outlined" />
          {article?.tags?.map((tag) => (
            <Chip key={tag} label={tag} size="small" variant="outlined" />
          ))}
        </Stack>

        <Stack direction="row" spacing={3} alignItems="center">
          <Typography variant="caption" color="text.secondary">
            📅 {article?.created_at}
          </Typography>

          <Box display="flex" alignItems="center" gap={0.5}>
            <VisibilityOutlinedIcon fontSize="small" />
            <Typography variant="caption">{article?.views}</Typography>
          </Box>

          <Box display="flex" alignItems="center" gap={0.5}>
            <FavoriteBorderOutlinedIcon fontSize="small" />
            <Typography variant="caption">{article?.likes}</Typography>
          </Box>

          <Box display="flex" alignItems="center" gap={0.5}>
            <ChatBubbleOutlineOutlinedIcon fontSize="small" />
            <Typography variant="caption">{article?.comments}</Typography>
          </Box>
        </Stack>
      </Stack>
    </div>
  );
};

export default ArticleHeader;
