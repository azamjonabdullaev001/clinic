/**
 * detectBtsBranch — matches a delivery address string against a list of
 * BTS branches loaded from the API (/api/bts-branches).
 * Returns the best-matching branch or null.
 */
export function detectBtsBranch(deliveryAddress, branches) {
  if (!deliveryAddress || !branches?.length) return null
  const lower = deliveryAddress.toLowerCase()
  // Try to match city/region name — longer names first to avoid false positives
  const sorted = [...branches].sort((a, b) => b.city.length - a.city.length)
  for (const branch of sorted) {
    const city = branch.city.toLowerCase()
    const region = branch.region?.toLowerCase() || ''
    if (lower.includes(city) || (region && lower.includes(region))) {
      return branch
    }
  }
  return null
}
