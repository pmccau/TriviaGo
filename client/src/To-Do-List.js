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
      categories: []
    };
  }

  componentDidMount() {
    // this.getQuestions();
    this.getCategories();
  }

  onChange = event => {
    this.setState({
      [event.target.name]: event.target.value
    });
  };

  onSubmit = () => {
    let { teamName } = this.state;
    if (teamName) {
      axios.post(endpoint + "/api/test",
          { teamName },
          { 
            headers: {
              "Content-Type": "application/x-www-form-urlencoded"
            }
          }
        ).then(res => {
          this.setState({ teamName: "" });
          console.log(res);
        });
    }
  };

  getCategories = () => {
    axios.get(endpoint + "/api/getCategories").then(res => {
      if (res.data) {
        console.log("CATEGORIES BELOW")
        console.log(res.data)
        // this.setState({
        //   categories: res.data
        // })
      }
    })
  }

  getQuestions = () => {
    axios.get(endpoint + "/api/getQuestions").then(res => {
      if (res.data) {
        this.setState({
          questions: res.data.map(question => {
            // return (<div className={"row"} key={question.Text}>{ question.Text }</div>
            // )
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

  // getTask = () => {
  //   axios.get(endpoint + "/api/task").then(res => {
  //     console.log(res);
  //     if (res.data) {
  //       this.setState({
  //         items: res.data.map(item => {
  //           let color = "yellow";

  //           if (item.status) {
  //             color = "green";
  //           }
  //           return (
  //             <Card key={item._id} color={color} fluid>
  //               <Card.Content>
  //                 <Card.Header textAlign="left">
  //                   <div style={{ wordWrap: "break-word" }}>{item.task}</div>
  //                 </Card.Header>

  //                 <Card.Meta textAlign="right">
  //                   <Icon
  //                     name="check circle"
  //                     color="green"
  //                     onClick={() => this.updateTask(item._id)}
  //                   />
  //                   <span style={{ paddingRight: 10 }}>Done</span>
  //                   <Icon
  //                     name="undo"
  //                     color="yellow"
  //                     onClick={() => this.undoTask(item._id)}
  //                   />
  //                   <span style={{ paddingRight: 10 }}>Undo</span>
  //                   <Icon
  //                     name="delete"
  //                     color="red"
  //                     onClick={() => this.deleteTask(item._id)}
  //                   />
  //                   <span style={{ paddingRight: 10 }}>Delete</span>
  //                 </Card.Meta>
  //               </Card.Content>
  //             </Card>
  //           );
  //         })
  //       });
  //     } else {
  //       this.setState({
  //         items: []
  //       });
  //     }
  //   });
  // };

  // updateTask = id => {
  //   axios
  //     .put(endpoint + "/api/task/" + id, {
  //       headers: {
  //         "Content-Type": "application/x-www-form-urlencoded"
  //       }
  //     })
  //     .then(res => {
  //       console.log(res);
  //       this.getTask();
  //     });
  // };

  // undoTask = id => {
  //   axios
  //     .put(endpoint + "/api/undoTask/" + id, {
  //       headers: {
  //         "Content-Type": "application/x-www-form-urlencoded"
  //       }
  //     })
  //     .then(res => {
  //       console.log(res);
  //       this.getTask();
  //     });
  // };

  // deleteTask = id => {
  //   axios
  //     .delete(endpoint + "/api/deleteTask/" + id, {
  //       headers: {
  //         "Content-Type": "application/x-www-form-urlencoded"
  //       }
  //     })
  //     .then(res => {
  //       console.log(res);
  //       this.getTask();
  //     });
  // };
  render() {
    return (
      <div>
        <div className="row">
          <Header className="header" as="h2">
            trivia.go
          </Header>
        </div>
        <div className="row">
          <Form onSubmit={this.onSubmit}>
            <Input
              type="text"
              name="teamName"
              onChange={this.onChange}
              value={this.state.teamName}
              fluid
              placeholder="Enter a team name..."
            />
            {/* <Button >Create teamName</Button> */}
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
