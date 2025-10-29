import ArticleFilterPanel from "@/components/ArticleFilterPanel";
import "./HomePage.scss";
import ArticleList from "@/components/ArticleList";
import TagCloud from "@/components/widgets/TagCloud";
import TwoColumnLayout from "@/components/layout/TwoColumnLayout";
import Aside from "@/components/Aside";
import HomeCommentList from "@/components/HomeCommentList";
import HomeBanner from "@/components/HomeBanner";

const HomePage = () => {
  return (
    <div className="home-page">
      <TwoColumnLayout
        main={
          <>
            <HomeBanner></HomeBanner>
            <ArticleFilterPanel />
            <ArticleList />
          </>
        }
        aside={
          <Aside>
            <TagCloud />
            <HomeCommentList></HomeCommentList>
          </Aside>
        }
      />
    </div>
  );
};

export default HomePage;
