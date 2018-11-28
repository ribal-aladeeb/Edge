import React, { Component } from "react";
import Post from "./post";
import PostEditor from "./postEditor";

class PostDisplay extends Component {
  constructor(props) {
    super(props);
    this.state = {
      email: this.props.email,
      posts: [],
      items: []
    };
  }

  addPost = newPostBody => {
    window.location.reload();
  };

  render() {
    return (
      <React.Fragment>
        <div className="card">
          <PostEditor addPost={this.addPost} email={this.props.email} />
        </div>

        <div className="postEdit">
          {this.props.posts.reverse().map((item, id) => {
            return (
              <Post
                id={id}
                key={id}
                postBody={item.data}
                from={item.author_email}
                posts={this.props.posts}
                postId={item.id}
                comments={this.props.comments}
                likes={this.props.likes}
                dislikes={this.props.dislikes}
                by={item.author_email}
                emailOfPost={this.props.email}
              />
            );
          })}
        </div>
      </React.Fragment>
    );
  }
}

export default PostDisplay;
