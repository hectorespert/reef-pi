import React from 'react'
import JournalForm from './form'
import PropTypes from 'prop-types'
import { fetchJournalUsage } from '../redux/actions/journal'
import { connect } from 'react-redux'

class Journal extends React.Component {
  constructor (props) {
    super(props)
    this.handleOnSubmit = this.handleOnSubmit.bind(this)
  }

  componentDidMount () {
    this.props.fetchJournalUsage(this.props.data.id)
  }

  handleOnSubmit (values) {
    this.props.onSubmit(values)
  }

  render () {
    if (this.props.readOnly) {
      if (this.props.journal_usage.data) { // TODO: Show last values
        return (
          <div>
            {this.props.journal_usage.data.current.map(function (d, idx) {
              return (<p key={idx}>{d.value}</p>)
            })}
          </div>
        )
      }
      return null
    } else {
      return (
        <JournalForm
          data={this.props.data}
          onSubmit={this.handleOnSubmit}
        />
      )
    }
  }
}

Journal.propTypes = {
  onSubmit: PropTypes.func.isRequired
}

const mapStateToProps = state => {
  return {
    journal_usage: state.journal_usage
  }
}

const mapDispatchToProps = dispatch => {
  return {
    fetchJournalUsage: (id) => dispatch(fetchJournalUsage(id))
  }
}

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(Journal)
