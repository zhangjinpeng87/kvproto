use eraftpb;
use pdpb;

impl From<pdpb::ConfChangeType> for eraftpb::ConfChangeType {
    fn from(ct: pdpb::ConfChangeType) -> Self {
        match ct {
            pdpb::ConfChangeType::AddNode => eraftpb::ConfChangeType::AddNode,
            pdpb::ConfChangeType::RemoveNode => eraftpb::ConfChangeType::RemoveNode,
            pdpb::ConfChangeType::UpdateNode => eraftpb::ConfChangeType::UpdateNode,
            pdpb::ConfChangeType::AddLearner => eraftpb::ConfChangeType::AddLearner,
        }
    }
}

impl From<eraftpb::ConfChangeType> for pdpb::ConfChangeType {
    fn from(ct: eraftpb::ConfChangeType) -> Self {
        match ct {
            eraftpb::ConfChangeType::AddNode => pdpb::ConfChangeType::AddNode,
            eraftpb::ConfChangeType::RemoveNode => pdpb::ConfChangeType::RemoveNode,
            eraftpb::ConfChangeType::UpdateNode => pdpb::ConfChangeType::UpdateNode,
            eraftpb::ConfChangeType::AddLearner => pdpb::ConfChangeType::AddLearner,
        }
    }
}
