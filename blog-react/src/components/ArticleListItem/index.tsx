import type { Article } from "@/api/article";
import { Visibility, Favorite, ChatBubble } from "@mui/icons-material";
import {
  Card,
  CardContent,
  CardHeader,
  Stack,
  Typography,
  Button,
} from "@mui/material";
import ArticleCover from "../widgets/ArticleCover";
import "./index.scss";
import { useNavigate } from "react-router-dom";

interface Props {
  id: string;
  article: Article;
}

const ArticleListItem = ({ id, article }: Props) => {
  const navigate = useNavigate();

  const handleReadMoreClick = () => {
    navigate(`/article/${id}`);
  };

  return (
    <Card className="article-list-item" variant="outlined">
      <CardHeader
        title={<Typography variant="h6">{article.title}</Typography>}
        subheader={article.created_at}
        action={
          <Stack direction="row" spacing={2}>
            <Typography variant="caption">
              <Visibility fontSize="inherit" /> {article.views}
            </Typography>
            <Typography variant="caption">
              <Favorite fontSize="inherit" /> {article.likes}
            </Typography>
            <Typography variant="caption">
              <ChatBubble fontSize="inherit" />
              {article.comments}
            </Typography>
          </Stack>
        }
      />

      <CardContent>
        <Stack direction="row" spacing={3}>
          <ArticleCover
            cover={`${import.meta.env.VITE_SERVER_URL}${article.cover}`}
          />

          <Stack spacing={2} flex={1}>
            <Typography variant="body2" color="text.secondary">
              {article.abstract}
            </Typography>
          </Stack>
          <Button
            size="small"
            className="article-readmore-btn"
            sx={{ alignSelf: "flex-end", whiteSpace: "nowrap" }}
            onClick={handleReadMoreClick}
          >
            READ MORE
          </Button>
        </Stack>
      </CardContent>
    </Card>
  );
};

export default ArticleListItem;
