// Code generated by entc, DO NOT EDIT.

package certificate

import (
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/mmildd_s/app/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Certificate {
	return predicate.Certificate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Certificate {
	return predicate.Certificate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Certificate {
	return predicate.Certificate(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Certificate {
	return predicate.Certificate(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Certificate {
	return predicate.Certificate(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Certificate {
	return predicate.Certificate(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Certificate {
	return predicate.Certificate(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Certificate {
	return predicate.Certificate(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Certificate {
	return predicate.Certificate(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CertificateName applies equality check predicate on the "Certificate_Name" field. It's identical to CertificateNameEQ.
func CertificateName(v string) predicate.Certificate {
	return predicate.Certificate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCertificateName), v))
	})
}

// CertificateNameEQ applies the EQ predicate on the "Certificate_Name" field.
func CertificateNameEQ(v string) predicate.Certificate {
	return predicate.Certificate(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCertificateName), v))
	})
}

// CertificateNameNEQ applies the NEQ predicate on the "Certificate_Name" field.
func CertificateNameNEQ(v string) predicate.Certificate {
	return predicate.Certificate(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCertificateName), v))
	})
}

// CertificateNameIn applies the In predicate on the "Certificate_Name" field.
func CertificateNameIn(vs ...string) predicate.Certificate {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Certificate(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCertificateName), v...))
	})
}

// CertificateNameNotIn applies the NotIn predicate on the "Certificate_Name" field.
func CertificateNameNotIn(vs ...string) predicate.Certificate {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Certificate(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCertificateName), v...))
	})
}

// CertificateNameGT applies the GT predicate on the "Certificate_Name" field.
func CertificateNameGT(v string) predicate.Certificate {
	return predicate.Certificate(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCertificateName), v))
	})
}

// CertificateNameGTE applies the GTE predicate on the "Certificate_Name" field.
func CertificateNameGTE(v string) predicate.Certificate {
	return predicate.Certificate(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCertificateName), v))
	})
}

// CertificateNameLT applies the LT predicate on the "Certificate_Name" field.
func CertificateNameLT(v string) predicate.Certificate {
	return predicate.Certificate(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCertificateName), v))
	})
}

// CertificateNameLTE applies the LTE predicate on the "Certificate_Name" field.
func CertificateNameLTE(v string) predicate.Certificate {
	return predicate.Certificate(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCertificateName), v))
	})
}

// CertificateNameContains applies the Contains predicate on the "Certificate_Name" field.
func CertificateNameContains(v string) predicate.Certificate {
	return predicate.Certificate(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldCertificateName), v))
	})
}

// CertificateNameHasPrefix applies the HasPrefix predicate on the "Certificate_Name" field.
func CertificateNameHasPrefix(v string) predicate.Certificate {
	return predicate.Certificate(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldCertificateName), v))
	})
}

// CertificateNameHasSuffix applies the HasSuffix predicate on the "Certificate_Name" field.
func CertificateNameHasSuffix(v string) predicate.Certificate {
	return predicate.Certificate(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldCertificateName), v))
	})
}

// CertificateNameEqualFold applies the EqualFold predicate on the "Certificate_Name" field.
func CertificateNameEqualFold(v string) predicate.Certificate {
	return predicate.Certificate(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldCertificateName), v))
	})
}

// CertificateNameContainsFold applies the ContainsFold predicate on the "Certificate_Name" field.
func CertificateNameContainsFold(v string) predicate.Certificate {
	return predicate.Certificate(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldCertificateName), v))
	})
}

// HasCertificateCoveredPerson applies the HasEdge predicate on the "Certificate_CoveredPerson" edge.
func HasCertificateCoveredPerson() predicate.Certificate {
	return predicate.Certificate(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CertificateCoveredPersonTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, CertificateCoveredPersonTable, CertificateCoveredPersonColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCertificateCoveredPersonWith applies the HasEdge predicate on the "Certificate_CoveredPerson" edge with a given conditions (other predicates).
func HasCertificateCoveredPersonWith(preds ...predicate.CoveredPerson) predicate.Certificate {
	return predicate.Certificate(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CertificateCoveredPersonInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, CertificateCoveredPersonTable, CertificateCoveredPersonColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Certificate) predicate.Certificate {
	return predicate.Certificate(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Certificate) predicate.Certificate {
	return predicate.Certificate(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Certificate) predicate.Certificate {
	return predicate.Certificate(func(s *sql.Selector) {
		p(s.Not())
	})
}
