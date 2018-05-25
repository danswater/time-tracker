import React, { Component } from 'react';
import { Container, Header, Segment, Grid, Button } from 'semantic-ui-react';
import {TimeContext } from './time-context';
import io from 'socket.io-client';


class App extends Component {

	componentDidMount () {
		const socket = io( 'localhost:8080//socket.io/' );
		console.log( socket );
	}

	renderTimer ( { state } ) {
		return (
			<React.Fragment>
				<Grid.Column width={4}>
					<Header as="h1">{state.Days}</Header>
					Days
				</Grid.Column>
				<Grid.Column width={4}>
					<Header as="h1">{state.Hours}</Header>
					Hours
				</Grid.Column>
				<Grid.Column width={4}>
					<Header as="h1">{state.Minutes}</Header>
					Minutes
				</Grid.Column>
				<Grid.Column width={4}>
					<Header as="h1">{state.Seconds}</Header>
					Seconds
				</Grid.Column>
			</React.Fragment>
		);
	}

	renderButtons ( { state, actions } ) {
		return (
			<Button.Group>
				<Button onClick={actions.startButton} color="green" disabled={!state.StartActive}>Start</Button>
				<Button onClick={actions.stopButton} color="red" disabled={!state.StopActive}>Stop</Button>
			</Button.Group>
		);
	}

	render() {
		return(
			<Container text style={{ marginTop: '2em' }} textAlign="center">
				<Header as="h1">Time Tracker</Header>
				<p>Synchronized timer</p>
				<Segment>
					<Grid>
						<TimeContext.Consumer>
							{this.renderTimer}
						</TimeContext.Consumer>
					</Grid>
				</Segment>
				<Grid>
					<Grid.Column width={16}>
						<TimeContext.Consumer>
							{this.renderButtons}
						</TimeContext.Consumer>
					</Grid.Column>
				</Grid>
			</Container>
		)
	}
}

export default App;
