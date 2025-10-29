import { Avatar, Box, Stack, Typography } from "@mui/material";
import { useNavigate } from "react-router-dom";
import "./index.scss";
import { useCommentNew } from "./hook/useCommentNew";

const HomeCommentList = () => {
  const navigate = useNavigate();
  const { comments } = useCommentNew();

  return (
    <Box className="home-comment-list">
      <Typography className="home-comment-title">Recent Comments</Typography>

      <Stack spacing={2}>
        {comments.map((comment) => (
          <Box key={comment.id} className="comment-item">
            <Box className="comment-header">
              <Box className="comment-user">
                <Avatar src={comment.user.avatar} className="comment-avatar" />
                <Typography className="comment-username">
                  {comment.user.username}
                </Typography>
              </Box>

              <Typography className="comment-time">
                {comment.created_at_text}
              </Typography>
            </Box>

            <Typography className="comment-content">
              {comment.content}
            </Typography>

            <Typography
              className="comment-action"
              onClick={() => navigate(`/article/${comment.article_id}`)}
            >
              查看文章 →
            </Typography>
          </Box>
        ))}
      </Stack>
    </Box>
  );
};

export default HomeCommentList;
