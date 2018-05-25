import React from 'react';

export const state = {
	Days        : 0,
	Hours       : 0,
	Minutes     : 0,
	Seconds     : 0,
	StopActive  : false,
	StartActive : true
};

export const actions = {};

export const TimeContext = React.createContext( {
	state,
	actions
} );
