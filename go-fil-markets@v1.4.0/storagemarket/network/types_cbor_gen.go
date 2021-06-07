// Code generated by github.com/whyrusleeping/cbor-gen. DO NOT EDIT.

package network

import (
	"fmt"
	"io"
	"sort"

	storagemarket "github.com/filecoin-project/go-fil-markets/storagemarket"
	crypto "github.com/filecoin-project/go-state-types/crypto"
	market "github.com/filecoin-project/specs-actors/actors/builtin/market"
	cid "github.com/ipfs/go-cid"
	cbg "github.com/whyrusleeping/cbor-gen"
	xerrors "golang.org/x/xerrors"
)

var _ = xerrors.Errorf
var _ = cid.Undef
var _ = sort.Sort

func (t *AskRequest) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{161}); err != nil {
		return err
	}

	scratch := make([]byte, 9)

	// t.Miner (address.Address) (struct)
	if len("Miner") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"Miner\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("Miner"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("Miner")); err != nil {
		return err
	}

	if err := t.Miner.MarshalCBOR(w); err != nil {
		return err
	}
	return nil
}

func (t *AskRequest) UnmarshalCBOR(r io.Reader) error {
	*t = AskRequest{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajMap {
		return fmt.Errorf("cbor input should be of type map")
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("AskRequest: map struct too large (%d)", extra)
	}

	var name string
	n := extra

	for i := uint64(0); i < n; i++ {

		{
			sval, err := cbg.ReadStringBuf(br, scratch)
			if err != nil {
				return err
			}

			name = string(sval)
		}

		switch name {
		// t.Miner (address.Address) (struct)
		case "Miner":

			{

				if err := t.Miner.UnmarshalCBOR(br); err != nil {
					return xerrors.Errorf("unmarshaling t.Miner: %w", err)
				}

			}

		default:
			// Field doesn't exist on this type, so ignore it
			cbg.ScanForLinks(r, func(cid.Cid) {})
		}
	}

	return nil
}
func (t *AskResponse) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{161}); err != nil {
		return err
	}

	scratch := make([]byte, 9)

	// t.Ask (storagemarket.SignedStorageAsk) (struct)
	if len("Ask") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"Ask\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("Ask"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("Ask")); err != nil {
		return err
	}

	if err := t.Ask.MarshalCBOR(w); err != nil {
		return err
	}
	return nil
}

func (t *AskResponse) UnmarshalCBOR(r io.Reader) error {
	*t = AskResponse{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajMap {
		return fmt.Errorf("cbor input should be of type map")
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("AskResponse: map struct too large (%d)", extra)
	}

	var name string
	n := extra

	for i := uint64(0); i < n; i++ {

		{
			sval, err := cbg.ReadStringBuf(br, scratch)
			if err != nil {
				return err
			}

			name = string(sval)
		}

		switch name {
		// t.Ask (storagemarket.SignedStorageAsk) (struct)
		case "Ask":

			{

				b, err := br.ReadByte()
				if err != nil {
					return err
				}
				if b != cbg.CborNull[0] {
					if err := br.UnreadByte(); err != nil {
						return err
					}
					t.Ask = new(storagemarket.SignedStorageAsk)
					if err := t.Ask.UnmarshalCBOR(br); err != nil {
						return xerrors.Errorf("unmarshaling t.Ask pointer: %w", err)
					}
				}

			}

		default:
			// Field doesn't exist on this type, so ignore it
			cbg.ScanForLinks(r, func(cid.Cid) {})
		}
	}

	return nil
}
func (t *Proposal) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{163}); err != nil {
		return err
	}

	scratch := make([]byte, 9)

	// t.DealProposal (market.ClientDealProposal) (struct)
	if len("DealProposal") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"DealProposal\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("DealProposal"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("DealProposal")); err != nil {
		return err
	}

	if err := t.DealProposal.MarshalCBOR(w); err != nil {
		return err
	}

	// t.Piece (storagemarket.DataRef) (struct)
	if len("Piece") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"Piece\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("Piece"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("Piece")); err != nil {
		return err
	}

	if err := t.Piece.MarshalCBOR(w); err != nil {
		return err
	}

	// t.FastRetrieval (bool) (bool)
	if len("FastRetrieval") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"FastRetrieval\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("FastRetrieval"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("FastRetrieval")); err != nil {
		return err
	}

	if err := cbg.WriteBool(w, t.FastRetrieval); err != nil {
		return err
	}
	return nil
}

func (t *Proposal) UnmarshalCBOR(r io.Reader) error {
	*t = Proposal{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajMap {
		return fmt.Errorf("cbor input should be of type map")
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("Proposal: map struct too large (%d)", extra)
	}

	var name string
	n := extra

	for i := uint64(0); i < n; i++ {

		{
			sval, err := cbg.ReadStringBuf(br, scratch)
			if err != nil {
				return err
			}

			name = string(sval)
		}

		switch name {
		// t.DealProposal (market.ClientDealProposal) (struct)
		case "DealProposal":

			{

				b, err := br.ReadByte()
				if err != nil {
					return err
				}
				if b != cbg.CborNull[0] {
					if err := br.UnreadByte(); err != nil {
						return err
					}
					t.DealProposal = new(market.ClientDealProposal)
					if err := t.DealProposal.UnmarshalCBOR(br); err != nil {
						return xerrors.Errorf("unmarshaling t.DealProposal pointer: %w", err)
					}
				}

			}
			// t.Piece (storagemarket.DataRef) (struct)
		case "Piece":

			{

				b, err := br.ReadByte()
				if err != nil {
					return err
				}
				if b != cbg.CborNull[0] {
					if err := br.UnreadByte(); err != nil {
						return err
					}
					t.Piece = new(storagemarket.DataRef)
					if err := t.Piece.UnmarshalCBOR(br); err != nil {
						return xerrors.Errorf("unmarshaling t.Piece pointer: %w", err)
					}
				}

			}
			// t.FastRetrieval (bool) (bool)
		case "FastRetrieval":

			maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
			if err != nil {
				return err
			}
			if maj != cbg.MajOther {
				return fmt.Errorf("booleans must be major type 7")
			}
			switch extra {
			case 20:
				t.FastRetrieval = false
			case 21:
				t.FastRetrieval = true
			default:
				return fmt.Errorf("booleans are either major type 7, value 20 or 21 (got %d)", extra)
			}

		default:
			// Field doesn't exist on this type, so ignore it
			cbg.ScanForLinks(r, func(cid.Cid) {})
		}
	}

	return nil
}
func (t *Response) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{164}); err != nil {
		return err
	}

	scratch := make([]byte, 9)

	// t.State (uint64) (uint64)
	if len("State") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"State\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("State"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("State")); err != nil {
		return err
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajUnsignedInt, uint64(t.State)); err != nil {
		return err
	}

	// t.Message (string) (string)
	if len("Message") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"Message\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("Message"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("Message")); err != nil {
		return err
	}

	if len(t.Message) > cbg.MaxLength {
		return xerrors.Errorf("Value in field t.Message was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len(t.Message))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string(t.Message)); err != nil {
		return err
	}

	// t.Proposal (cid.Cid) (struct)
	if len("Proposal") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"Proposal\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("Proposal"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("Proposal")); err != nil {
		return err
	}

	if err := cbg.WriteCidBuf(scratch, w, t.Proposal); err != nil {
		return xerrors.Errorf("failed to write cid field t.Proposal: %w", err)
	}

	// t.PublishMessage (cid.Cid) (struct)
	if len("PublishMessage") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"PublishMessage\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("PublishMessage"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("PublishMessage")); err != nil {
		return err
	}

	if t.PublishMessage == nil {
		if _, err := w.Write(cbg.CborNull); err != nil {
			return err
		}
	} else {
		if err := cbg.WriteCidBuf(scratch, w, *t.PublishMessage); err != nil {
			return xerrors.Errorf("failed to write cid field t.PublishMessage: %w", err)
		}
	}

	return nil
}

func (t *Response) UnmarshalCBOR(r io.Reader) error {
	*t = Response{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajMap {
		return fmt.Errorf("cbor input should be of type map")
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("Response: map struct too large (%d)", extra)
	}

	var name string
	n := extra

	for i := uint64(0); i < n; i++ {

		{
			sval, err := cbg.ReadStringBuf(br, scratch)
			if err != nil {
				return err
			}

			name = string(sval)
		}

		switch name {
		// t.State (uint64) (uint64)
		case "State":

			{

				maj, extra, err = cbg.CborReadHeaderBuf(br, scratch)
				if err != nil {
					return err
				}
				if maj != cbg.MajUnsignedInt {
					return fmt.Errorf("wrong type for uint64 field")
				}
				t.State = uint64(extra)

			}
			// t.Message (string) (string)
		case "Message":

			{
				sval, err := cbg.ReadStringBuf(br, scratch)
				if err != nil {
					return err
				}

				t.Message = string(sval)
			}
			// t.Proposal (cid.Cid) (struct)
		case "Proposal":

			{

				c, err := cbg.ReadCid(br)
				if err != nil {
					return xerrors.Errorf("failed to read cid field t.Proposal: %w", err)
				}

				t.Proposal = c

			}
			// t.PublishMessage (cid.Cid) (struct)
		case "PublishMessage":

			{

				b, err := br.ReadByte()
				if err != nil {
					return err
				}
				if b != cbg.CborNull[0] {
					if err := br.UnreadByte(); err != nil {
						return err
					}

					c, err := cbg.ReadCid(br)
					if err != nil {
						return xerrors.Errorf("failed to read cid field t.PublishMessage: %w", err)
					}

					t.PublishMessage = &c
				}

			}

		default:
			// Field doesn't exist on this type, so ignore it
			cbg.ScanForLinks(r, func(cid.Cid) {})
		}
	}

	return nil
}
func (t *SignedResponse) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{162}); err != nil {
		return err
	}

	scratch := make([]byte, 9)

	// t.Response (network.Response) (struct)
	if len("Response") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"Response\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("Response"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("Response")); err != nil {
		return err
	}

	if err := t.Response.MarshalCBOR(w); err != nil {
		return err
	}

	// t.Signature (crypto.Signature) (struct)
	if len("Signature") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"Signature\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("Signature"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("Signature")); err != nil {
		return err
	}

	if err := t.Signature.MarshalCBOR(w); err != nil {
		return err
	}
	return nil
}

func (t *SignedResponse) UnmarshalCBOR(r io.Reader) error {
	*t = SignedResponse{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajMap {
		return fmt.Errorf("cbor input should be of type map")
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("SignedResponse: map struct too large (%d)", extra)
	}

	var name string
	n := extra

	for i := uint64(0); i < n; i++ {

		{
			sval, err := cbg.ReadStringBuf(br, scratch)
			if err != nil {
				return err
			}

			name = string(sval)
		}

		switch name {
		// t.Response (network.Response) (struct)
		case "Response":

			{

				if err := t.Response.UnmarshalCBOR(br); err != nil {
					return xerrors.Errorf("unmarshaling t.Response: %w", err)
				}

			}
			// t.Signature (crypto.Signature) (struct)
		case "Signature":

			{

				b, err := br.ReadByte()
				if err != nil {
					return err
				}
				if b != cbg.CborNull[0] {
					if err := br.UnreadByte(); err != nil {
						return err
					}
					t.Signature = new(crypto.Signature)
					if err := t.Signature.UnmarshalCBOR(br); err != nil {
						return xerrors.Errorf("unmarshaling t.Signature pointer: %w", err)
					}
				}

			}

		default:
			// Field doesn't exist on this type, so ignore it
			cbg.ScanForLinks(r, func(cid.Cid) {})
		}
	}

	return nil
}
func (t *DealStatusRequest) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{162}); err != nil {
		return err
	}

	scratch := make([]byte, 9)

	// t.Proposal (cid.Cid) (struct)
	if len("Proposal") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"Proposal\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("Proposal"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("Proposal")); err != nil {
		return err
	}

	if err := cbg.WriteCidBuf(scratch, w, t.Proposal); err != nil {
		return xerrors.Errorf("failed to write cid field t.Proposal: %w", err)
	}

	// t.Signature (crypto.Signature) (struct)
	if len("Signature") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"Signature\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("Signature"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("Signature")); err != nil {
		return err
	}

	if err := t.Signature.MarshalCBOR(w); err != nil {
		return err
	}
	return nil
}

func (t *DealStatusRequest) UnmarshalCBOR(r io.Reader) error {
	*t = DealStatusRequest{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajMap {
		return fmt.Errorf("cbor input should be of type map")
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("DealStatusRequest: map struct too large (%d)", extra)
	}

	var name string
	n := extra

	for i := uint64(0); i < n; i++ {

		{
			sval, err := cbg.ReadStringBuf(br, scratch)
			if err != nil {
				return err
			}

			name = string(sval)
		}

		switch name {
		// t.Proposal (cid.Cid) (struct)
		case "Proposal":

			{

				c, err := cbg.ReadCid(br)
				if err != nil {
					return xerrors.Errorf("failed to read cid field t.Proposal: %w", err)
				}

				t.Proposal = c

			}
			// t.Signature (crypto.Signature) (struct)
		case "Signature":

			{

				if err := t.Signature.UnmarshalCBOR(br); err != nil {
					return xerrors.Errorf("unmarshaling t.Signature: %w", err)
				}

			}

		default:
			// Field doesn't exist on this type, so ignore it
			cbg.ScanForLinks(r, func(cid.Cid) {})
		}
	}

	return nil
}
func (t *DealStatusResponse) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{162}); err != nil {
		return err
	}

	scratch := make([]byte, 9)

	// t.DealState (storagemarket.ProviderDealState) (struct)
	if len("DealState") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"DealState\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("DealState"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("DealState")); err != nil {
		return err
	}

	if err := t.DealState.MarshalCBOR(w); err != nil {
		return err
	}

	// t.Signature (crypto.Signature) (struct)
	if len("Signature") > cbg.MaxLength {
		return xerrors.Errorf("Value in field \"Signature\" was too long")
	}

	if err := cbg.WriteMajorTypeHeaderBuf(scratch, w, cbg.MajTextString, uint64(len("Signature"))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string("Signature")); err != nil {
		return err
	}

	if err := t.Signature.MarshalCBOR(w); err != nil {
		return err
	}
	return nil
}

func (t *DealStatusResponse) UnmarshalCBOR(r io.Reader) error {
	*t = DealStatusResponse{}

	br := cbg.GetPeeker(r)
	scratch := make([]byte, 8)

	maj, extra, err := cbg.CborReadHeaderBuf(br, scratch)
	if err != nil {
		return err
	}
	if maj != cbg.MajMap {
		return fmt.Errorf("cbor input should be of type map")
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("DealStatusResponse: map struct too large (%d)", extra)
	}

	var name string
	n := extra

	for i := uint64(0); i < n; i++ {

		{
			sval, err := cbg.ReadStringBuf(br, scratch)
			if err != nil {
				return err
			}

			name = string(sval)
		}

		switch name {
		// t.DealState (storagemarket.ProviderDealState) (struct)
		case "DealState":

			{

				if err := t.DealState.UnmarshalCBOR(br); err != nil {
					return xerrors.Errorf("unmarshaling t.DealState: %w", err)
				}

			}
			// t.Signature (crypto.Signature) (struct)
		case "Signature":

			{

				if err := t.Signature.UnmarshalCBOR(br); err != nil {
					return xerrors.Errorf("unmarshaling t.Signature: %w", err)
				}

			}

		default:
			// Field doesn't exist on this type, so ignore it
			cbg.ScanForLinks(r, func(cid.Cid) {})
		}
	}

	return nil
}