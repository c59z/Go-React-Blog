import type { Advertisement } from "@/api/advertisement";
import "./index.scss";

interface Props {
  item: Advertisement;
}

const AdvertisementSlide = ({ item }: Props) => {
  return (
    <a className="ad-slide" href={item.link} target="_blank" rel="noreferrer">
      <img
        src={`${import.meta.env.VITE_SERVER_URL}${item.ad_image}`}
        alt={item.title}
      />
      <div className="ad-mask">
        <h3>{item.title}</h3>
        <p>{item.content}</p>
      </div>
    </a>
  );
};

export default AdvertisementSlide;
