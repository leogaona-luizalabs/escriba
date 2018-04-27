package draft

const insertSQL = `insert into drafts (url) values (?)`
const approveSQL = `update drafts d set approvals = d.approvals + 1 where url = ?`
const publishSQL = `update drafts d set published_at = NOW() where url = ?`
const listReviewSQL = `select url, approvals, created_at from drafts where published_at is null and approvals < ? order by created_at`
const listPublicationSQL = `select url, approvals, created_at from drafts where published_at is null and approvals >= ? order by created_at`
