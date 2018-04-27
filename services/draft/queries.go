package draft

const insertSQL = `insert into draft (url) values (?)`
const approveSQL = `update draft d set approvals = d.approvals + 1 where url = ?`
const publishSQL = `update draft d set published_at = NOW() where url = ?`
const listReviewSQL = `select url, approvals from draft where published_at is null and approvals < ?`
const listPublicationSQL = `select url, approvals from draft where published_at is null and approvals >= ?`
