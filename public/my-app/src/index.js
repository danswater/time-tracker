import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import App from './App';
import registerServiceWorker from './registerServiceWorker';
import 'semantic-ui-css/semantic.min.css';

import {TimeContext, state } from './time-context';

class Provider extends React.Component {
	constructor( props ) {
		super(props);
		this.state = { state, actions : {
			startButton : this.startButton.bind( this ),
			stopButton  : this.stopButton.bind( this )
		} };
	}

	startButton () {
		const state = Object.assign( {}, this.state.state, {
			StopActive  : true,
			StartActive : false
		} );
		this.setState( Object.assign( {}, this.state, { state } ) );
	}

	stopButton () {
		const state = Object.assign( {}, this.state.state, {
			StopActive  : false,
			StartActive : true
		} );
		this.setState( Object.assign( {}, this.state, { state } ) );
	}

	render () {
		return (
			<TimeContext.Provider value={this.state}>
				{this.props.children}
			</TimeContext.Provider>
		);
	}
}


ReactDOM.render(
	<Provider>
		<App />
	</Provider>,
	document.getElementById('root')
);
registerServiceWorker();
