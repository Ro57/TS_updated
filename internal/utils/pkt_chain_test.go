package utils

// Make a test dummy implementation MockPktChain of this where:
// CurrentHeight() -> (unixTime() - 1566269808) / 60
// BlockHashAtHeight(height) -> if height > CurrentHeight() { nil } else { sha256(height) }
// AnnounceData(data) -> go func() { loop { sleepSeconds(random(30, 90)); channel <- AnnProof { num: CurrentHeight() } } }
// VerifyProof(ap) -> return ap.num
