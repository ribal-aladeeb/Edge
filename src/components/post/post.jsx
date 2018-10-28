import React from "react";
import CommentDisplay from "./commentDisplay";
import "./css/post.css";

const Post = props => (
  <div className="aPost">
    <div className="card">
      <h5>Post by {localStorage.getItem("email")}</h5>
      {props.postBody}
      <br />
      <br />
      <CommentDisplay />
      <a href="/">Like Dislike</a>
    </div>
  </div>
);

export default Post;
