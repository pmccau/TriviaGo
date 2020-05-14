import React, { Component } from "react";
import axios from "axios";
import { Card, Header, Form, Input } from "semantic-ui-react";

let endpoint = "http://localhost:8080";

class ToDoList extends Component {
  constructor(props) {
    super(props);

    this.state = {
      teamName: "",
      items: [],
      questions: [],
      categories: [],
      categoryID: "",
      numQuestions: 10,
      questionType: "",
      difficulty: "",
      category: "",
    };
  }

  componentDidMount() {
    // this.getQuestions();
    this.getCategories();
  }

  onChange = event => {
    console.log([event.target.name] + ": " + event.target.value)
    this.setState({
      [event.target.name]: event.target.value
    });
  };

  onSubmit = () => {
    let { numQuestions } = this.state;
    let { difficulty } = this.state;
    let { category } = this.state;
    let { questionType } = this.state;
    console.log("numQuestions=" + numQuestions)
    if (numQuestions && difficulty && category && questionType) {
      axios.post(endpoint + "/api/test",
          { numQuestions, difficulty, category, questionType },
          { 
            headers: {
              "Content-Type": "application/x-www-form-urlencoded"
            }
          }
        ).then(res => {
          this.setState({ numQuestions: 10 });
          console.log(res);
        });
    }
  };

  getCategories = () => {
    axios.get(endpoint + "/api/getCategories").then(res => {
      if (res.data) {
        console.log("CATEGORIES BELOW")
        console.log(res.data)
        this.setState({
          categories: res.data.map(category => {
            return (<a
                key={ category.ID }
                value={ this.state.categoryID }
                name="category"
                onClick={ this.onChange }
                href="#">{ category.Name }</a>)
          })
        })
      }
    })
  }

  getQuestions = () => {
    axios.get(endpoint + "/api/getQuestions").then(res => {
      if (res.data) {
        this.setState({
          questions: res.data.map(question => {
            return (<div className={"row"} key={question.Text}>{ question.Text }</div>
            )
            let color = "green";
            if (question.Difficulty === "medium") {
              color = "yellow";
            } else if (question.Difficulty === "hard") {
              color = "red";
            }
            return (
              <Card key={question.Text} color={color} fluid>
                <Card.Content>
                  <Card.Header textAlign="left">
                    <div style={{ wordWrap: "break-word", textAlign: "left" }}>{question.Text}</div>
                  </Card.Header>
                </Card.Content>
              </Card>
            )
          })
        })
      }
    })
  }

  getCategoryDropdowns = () => {
    console.log(this.state.categories)
    var output = ""
    for (var i = 0; i < this.state.categories.length; i++) {
      output += <a href="#">{ this.state.categories[i].name }</a>
    }
    return output
  }


  render() {
    return (
      <div>
        <div className="row">
          <Header className="header" as="h2">
            trivia.go
          </Header>
        </div>

        <div className="navbar">
          <a href="#home">Home</a>
          <a href="#news">News</a>
          <div className="dropdown" value={"test1"}>
            <button className="dropbtn" onClick={this.onChange} value={"abcd"}>Categories
              <i className="fa fa-caret-down"></i>
            </button>
            <div className="dropdown-content" value={"test"}>
              {this.state.categories}
            </div>
          </div>
        </div>

        <div className="row">

          <Input
            type="number"
            name="numQuestions"
            onChange={this.onChange}
            value={this.state.numQuestions}
            fluid
            placeholder={10}
          />
          <Input
              type="text"
              name="category"
              onChange={this.onChange}
              value={this.state.category}
              fluid
              placeholder="category: ex., 'General Knowledge'"
          />
          <Input
              type="text"
              name="questionType"
              onChange={this.onChange}
              value={this.state.questionType}
              fluid
              placeholder="questionType: 'true/false', 'multiple choice'"
          />

        <Form onSubmit={this.onSubmit}>
          <Input
                type="text"
                name="difficulty"
                onChange={this.onChange}
                value={this.state.difficulty}
                fluid
                placeholder="Easy"
            />
        </Form>

        </div>
        <div className="row">
          <Card.Group>{this.state.questions}</Card.Group>
        </div>
      </div>
    );
  }
}

export default ToDoList;
